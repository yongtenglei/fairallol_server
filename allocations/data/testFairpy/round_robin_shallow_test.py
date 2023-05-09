import fairpy
import time
import os

divide = fairpy.divide

data_folder = "../"
agents = ["Alice", "Bob"]


def process_file(file_path, pattern):
    with open(file_path, "r") as file:
        lines = file.readlines()

    input_2_agents = {}
    if pattern == "tie":
        line = lines[0].strip()
        values = [int(value) for value in line.split(",")]
        agent_data = {}
        for j, value in enumerate(values, 1):
            agent_data[f"Item{j}"] = value
        for agent in agents:
            input_2_agents[agent] = agent_data
    else:
        for i, line in enumerate(lines):
            line = line.strip()
            values = [int(value) for value in line.split(",")]
            agent_data = {}
            for j, value in enumerate(values, 1):
                agent_data[f"Item{j}"] = value
            input_2_agents[agents[i]] = agent_data

    return input_2_agents


for file in os.listdir(data_folder):
    if (
        file.endswith("_shallow_normal.txt")
        or file.endswith("_shallow_similar.txt")
        or file.endswith("_shallow_tie.txt")
    ):
        file_path = os.path.join(data_folder, file)
        pattern = file.split("_")[2].split(".")[0]
        input_2_agents = process_file(file_path, pattern)

        start_time = time.time()  # Record the start time
        print(
            "{}: \n".format(pattern), divide(fairpy.items.round_robin, input_2_agents)
        )
        end_time = time.time()  # Record the end time
        elapsed_time = end_time - start_time  # Calculate the elapsed time
        print("elapsed time: ", elapsed_time)
        print()
