import fairpy
import time
import os


divide = fairpy.divide

data_folder = "../"
agents = ["Alice", "Bob"]


def process_file(file_path, pattern):
    with open(file_path, "r") as file:
        lines = file.readlines()

    input_2_agents_list = []

    if pattern == "tie":
        for line in lines:
            line = line.strip()
            values = [int(value) for value in line.split(",")]
            agent_data = {}
            for j, value in enumerate(values, 1):
                agent_data[f"Item{j}"] = value
            input_2_agents = {agent: agent_data for agent in agents}
            input_2_agents_list.append(input_2_agents)
    else:
        for i in range(0, len(lines), 2):
            input_2_agents = {}
            for agent_idx, line in enumerate(lines[i : i + 2]):
                line = line.strip()
                values = [int(value) for value in line.split(",")]
                agent_data = {}
                for j, value in enumerate(values, 1):
                    agent_data[f"Item{j}"] = value
                input_2_agents[agents[agent_idx]] = agent_data
            input_2_agents_list.append(input_2_agents)

    return input_2_agents_list


output_dict = {
    "normal": {4: [], 5: [], 6: []},
    "similar": {4: [], 5: [], 6: []},
    "tie": {4: [], 5: [], 6: []},
}

for file in os.listdir(data_folder):
    if "_data_" in file:
        file_path = os.path.join(data_folder, file)
        file_fields = file.split("_")
        N = int(file_fields[0])
        pattern = file_fields[2]

        input_2_agents_list = process_file(file_path, pattern)
        output_dict[pattern][N] = input_2_agents_list

# for pattern in output_dict.keys():
# for N in output_dict[pattern].keys():
# length = len(output_dict[pattern][N])
# print(f"Pattern: {pattern}, N: {N}, Length: {length}")

for pattern in output_dict.keys():
    start_time = time.time()  # Record the start time
    for N in output_dict[pattern].keys():
        for i in range(len(output_dict[pattern][N])):
            divide(fairpy.items.round_robin, output_dict[pattern][N][i])

        end_time = time.time()  # Record the end time
        elapsed_time = end_time - start_time  # Calculate the elapsed time
        print("Test for: ")
        print(f"\tN: {N}, Pattern: {pattern}, Time elapsed: {elapsed_time}")
