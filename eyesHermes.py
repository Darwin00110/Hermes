import pyautogui
import numpy as np
import cv2
import json
import os


pathJSON = os.path.realpath(
    os.path.join(os.getenv("APPDATA"), "..", "Local", "Hermes", "Data.json")
)
pathDirJSON = os.path.dirname(pathJSON)


def loadJSON():
    with open(pathJSON, "r") as file:
        return json.load(file)


def configProgram():
    if not os.path.exists(pathDirJSON):
        os.makedirs(pathDirJSON)
    if os.path.exists(pathJSON):
        return
    else:
        with open(pathJSON, "w") as file:
            file.close()


def verificar_imagem_na_tela(template_path, threshold=0.8):
    # 1. Capturar a tela
    screenshot = pyautogui.screenshot()

    # 2. Converter a captura para um formato compatível com OpenCV (BGR)
    screenshot = cv2.cvtColor(np.array(screenshot), cv2.COLOR_RGB2BGR)

    # 3. Carregar a imagem que você está procurando
    template = cv2.imread(template_path)
    if template is None:
        print("Erro: Imagem de referência não encontrada.")
        return False

    # 4. Template Matching (Procurar a imagem na tela)
    result = cv2.matchTemplate(screenshot, template, cv2.TM_CCOEFF_NORMED)

    # 5. Obter a melhor correspondência
    min_val, max_val, min_loc, max_loc = cv2.minMaxLoc(result)

    # 6. Verificar se a correspondência supera o limite (threshold)
    if max_val >= threshold:
        x = max_loc[0] + 18
        y = max_loc[1] + 12
        new_loc = (x, y)
        if not os.path.exists(pathJSON):
            Data = {
                "Taxa de aprovacao": max_val,
                "Localizacao": {
                    "x": new_loc[0],
                    "y": new_loc[1],
                },
                "Encontrado": True,
            }
            json.dump(Data, open(pathJSON, "w"), indent=4)
            return True
        else:
            with open(pathJSON, "r") as file:
                ContentJSON = json.load(file)
                ContentJSON["Localizacao"]["x"] = new_loc[0]
                ContentJSON["Localizacao"]["y"] = new_loc[1]
                file.close()
            with open(pathJSON, "w") as file:
                json.dump(ContentJSON, file, indent=4)

    else:
        Data = {"Taxa de aprovacao": max_val, "Localizacao": None, "Encontrado": False}
        print(f"Imagem não encontrada. Melhor chance: {max_val:.2f}")
        return False


# --- Exemplo de uso ---
# Certifique-se de ter um arquivo 'botao.png' na mesma pasta
def teste():
    pyautogui.moveTo(20, 98)
    # pyautogui.click()

def main():
    jsonPROGRAM = loadJSON()
    print(jsonPROGRAM["pathImageUser"])
    if verificar_imagem_na_tela(jsonPROGRAM["pathImageUser"]):
        print("Ação: Botão está visível!")
    else:
        print("Ação: Botão não está visível.")
configProgram()
main()
