import itertools
import random
import numpy as np


def validate(result):
    for sample in result:
        assert len(sample) == n, f"length is not equal to {n}"
        assert sum(sample) == total_sum, f"sum is not equal to {total_sum}"


def generate_random_data(n, total_sum, num_samples):
    """
    Generate `num_samples` of data,
    each with `n` evaluations,
    and sum to `total_sum`
    """

    # Generate a combination of n - 1 integers in the range 1 to total_sum - 1
    combinations = list(itertools.combinations(range(1, total_sum), n - 1))

    # shuffle data
    random.shuffle(combinations)

    selected_combinations = combinations[:num_samples]

    # Calculate the difference between adjacent elements
    # We got (2, 4, 7) -> (0, 2, 4, 7, 10)
    # Then
    # 2 - 0 = 2
    # 4 - 2 = 2
    # 7 - 4 = 3
    # 10 - 7 = 3
    # we got (2, 2, 3, 3) -> sum = 10, n = 4
    data = []
    for combination in selected_combinations:
        combination = (0,) + combination + (total_sum,)
        sample = [combination[i + 1] - combination[i] for i in range(n)]
        data.append(sample)

    return data


def generate_similar_random_data(n, total_sum, num_samples, similar):
    """
    Generate `num_samples` of data,
    each with `n` evaluations,
    and sum to `total_sum`,
    each `similar` data has a certain similarity.
    """

    # Generate a combination of n - 1 integers in the range 1 to total_sum - 1
    combinations = list(itertools.combinations(range(1, total_sum), n - 1))

    # Define the interval between adjacent combinations
    interval = len(combinations) // (num_samples * similar)

    # When selecting the first num_samples * similar combinations, randomly select similar combinations within a fixed interval each time
    selected_combinations = []
    for i in range(0, num_samples * interval * similar, interval * similar):
        group = [
            combinations[j]
            for j in random.sample(range(i, i + interval * similar), similar)
        ]
        selected_combinations.append(group)

    # Calculate the difference between adjacent elements
    data = []
    for group in selected_combinations:
        for combination in group:
            combination = (0,) + combination + (total_sum,)
            sample = [combination[i + 1] - combination[i] for i in range(n)]
            data.append(sample)

    return data


n_values = [4, 5, 6, 7]
patterns = ["normal", "similar", "tie"]
total_sum = 100
similar = 10

for n in n_values:
    for pattern in patterns:
        num_samples = 1000

        result = []

        if pattern == "tie":
            num_samples = 500
        elif pattern == "similar":
            num_samples = num_samples // similar
            print(f"n = {n} pattern = {pattern} num_samples = {num_samples}")
            # use generate_similar_random_data
            result = generate_similar_random_data(n, total_sum, num_samples, similar)
            validate(result)
            np.savetxt(
                f"{n}_data_{pattern}_{num_samples}.txt", result, fmt="%d", delimiter=","
            )
            print(f"save data to {n}_data_{pattern}_{num_samples}.txt")
            continue

        print(f"n = {n} pattern = {pattern} num_samples = {num_samples}")
        result = generate_random_data(n, total_sum, num_samples)
        validate(result)
        np.savetxt(
            f"{n}_data_{pattern}_{num_samples}.txt", result, fmt="%d", delimiter=","
        )
        print(f"save data to {n}_data_{pattern}_{num_samples}.txt")
