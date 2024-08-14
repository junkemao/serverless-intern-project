import numpy as np
import matplotlib.pyplot as plt
import os

def convert_file_to_numpy_array(file_path):
    numbers = []

    with open(file_path, 'r') as file:
        next(file)
        for line in file:
            parts = line.split()
            if parts:
                try:
                    number = float(parts[0])
                    numbers.append(number)
                except ValueError:
                    print(f"Warning: Could not convert '{parts[0]}' to float.")

    numbers_array = np.array(numbers)
    
    return numbers_array

def plot_cdf(numbers_array, output_file):
    sorted_numbers = np.sort(numbers_array)
    
    cdf = np.arange(1, len(sorted_numbers)+1) / len(sorted_numbers)
    
    plt.scatter(sorted_numbers, cdf, s=1)
    plt.xlabel('latency')
    plt.ylabel('CDF')
    plt.title('Duration=400ms | RPS=25')
    
    plt.savefig(output_file)
    plt.close()

file_path = input("Enter the path to the input file: ")
numbers_array = convert_file_to_numpy_array(file_path)
output_file = 'cdf_scatter_plot.png'
plot_cdf(numbers_array, output_file)

print(f"CDF scatter plot saved as {output_file}")
