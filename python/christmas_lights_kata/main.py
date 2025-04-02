# https://kata-log.rocks/christmas-lights-kata


class Light:
    def __init__(self):
        self.is_on = False

    def turn(self, on: bool):
        self.is_on = on

    def toggle(self):
        self.is_on = not self.is_on


class Instruction:
    TURN_ON = "turn_on"
    TURN_OFF = "turn_off"
    TOGGLE = "toggle"

    def __init__(self, instruction: str, start: tuple[int, int], end: tuple[int, int]):
        self.instruction = instruction
        self.start = start
        self.end = end


lights: list[list[Light]] = []


def init_lights():
    for i in range(1000):
        new_light_list: list[Light] = []
        lights.append(new_light_list)
        for j in range(1000):
            lights[i].append(Light())


init_lights()

instructions: list[Instruction] = [
    Instruction(Instruction.TURN_ON, (887, 9), (959, 629)),
    Instruction(Instruction.TURN_ON, (454, 398), (844, 448)),
    Instruction(Instruction.TURN_OFF, (539, 243), (559, 965)),
    Instruction(Instruction.TURN_OFF, (370, 819), (676, 868)),
    Instruction(Instruction.TURN_OFF, (145, 40), (370, 997)),
    Instruction(Instruction.TURN_OFF, (301, 3), (808, 453)),
    Instruction(Instruction.TURN_ON, (351, 678), (951, 908)),
    Instruction(Instruction.TOGGLE, (720, 196), (897, 994)),
    Instruction(Instruction.TOGGLE, (831, 394), (904, 860)),
]

for intruction in instructions:
    x_quantity = intruction.end[0] - intruction.start[0] + 1
    y_quantity = intruction.end[1] - intruction.start[1] + 1

    for x in range(x_quantity):
        for y in range(y_quantity):
            if intruction.instruction == Instruction.TURN_ON:
                lights[intruction.start[0] + x][intruction.start[1] + y].turn(True)
            elif intruction.instruction == Instruction.TURN_OFF:
                lights[intruction.start[0] + x][intruction.start[1] + y].turn(False)
            elif intruction.instruction == Instruction.TOGGLE:
                lights[intruction.start[0] + x][intruction.start[1] + y].toggle()

lights_on = 0
for light in lights:
    for light_row in light:
        if light_row.is_on:
            lights_on += 1

print(lights_on)
