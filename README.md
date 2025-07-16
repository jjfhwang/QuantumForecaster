# QuantumForecaster: Unveiling Forex Microstructure via Wavelet-Decomposed Order Flow

This repository contains a Go-based implementation for quantifying Forex market microstructure dynamics using wavelet analysis of order flow imbalances derived from tick data. The project facilitates the generation of latency-sensitive trading signals and allows for rigorous backtesting to assess signal efficacy.

This project provides a framework for researchers and quantitative traders to explore the subtle intricacies of Forex market microstructure. By leveraging wavelet decomposition, QuantumForecaster isolates and analyzes order flow imbalances across different time scales, effectively filtering out noise and revealing underlying patterns that might be obscured in raw tick data. The core objective is to translate these patterns into actionable trading signals. The system ingests high-frequency tick data, calculates order flow imbalances (buyer-initiated versus seller-initiated trades), performs wavelet decomposition on these imbalances, and then uses statistical features extracted from the wavelet coefficients to generate trading signals. A comprehensive backtesting module is included, enabling users to evaluate the performance of these signals using historical tick data. This backtesting engine simulates realistic trading conditions, accounting for latency and transaction costs.

QuantumForecaster goes beyond simple technical analysis by incorporating principles of signal processing and time-frequency analysis. Wavelet decomposition allows for the identification of short-lived and long-lived patterns in order flow, providing a more nuanced understanding of market dynamics than traditional methods. The backtesting module is designed to be highly configurable, allowing users to experiment with different signal generation strategies, risk management parameters, and trading styles. The project provides a flexible and extensible platform for developing and testing sophisticated trading algorithms based on market microstructure analysis. The goal is to provide a robust tool for researching and implementing latency-sensitive trading strategies in the Forex market.

The projects emphasis on high-frequency data and real-time signal generation makes it particularly relevant for participants operating in the arena of algorithmic trading. The provided backtesting framework provides valuable insights into the robustness of signal generation strategies prior to real-world deployment. Furthermore, the modular design of the system allows for easy integration with existing trading infrastructure. Researchers can use the project to explore novel trading strategies and gain a deeper understanding of Forex market dynamics.

## Key Features

*   **Tick Data Ingestion:** Efficiently processes high-frequency tick data from various sources, including CSV files and real-time data streams (e.g., via FIX protocol). This is achieved using Go's concurrency features to ensure minimal latency in data processing. A custom data structure optimizes storage and retrieval of tick data.
*   **Order Flow Imbalance Calculation:** Computes order flow imbalances based on bid-ask quotes and trade direction. The algorithm classifies trades as buyer-initiated or seller-initiated based on their proximity to the prevailing bid or ask price.
*   **Wavelet Decomposition:** Implements a discrete wavelet transform (DWT) using a Daubechies wavelet family (configurable). The decomposition is performed on the calculated order flow imbalances, separating the signal into different frequency bands.
*   **Signal Generation:** Generates trading signals based on statistical features extracted from the wavelet coefficients (e.g., variance, kurtosis, skewness of each scale). A rule-based system allows for the creation of custom signal generation logic based on these features.
*   **Backtesting Engine:** Provides a comprehensive backtesting engine that simulates trading activity using historical tick data. The backtesting engine accounts for latency, transaction costs, and slippage.
*   **Performance Metrics:** Calculates various performance metrics, including Sharpe ratio, maximum drawdown, and profit factor, to evaluate the efficacy of trading strategies.
*   **Configurable Parameters:** The system is highly configurable, allowing users to adjust parameters such as wavelet family, decomposition level, signal generation thresholds, and backtesting parameters.

## Technology Stack

*   **Go:** The primary programming language, chosen for its performance, concurrency support, and suitability for building high-frequency trading systems.
*   **gonum.org/v1/gonum:** A numerical computing library for Go, used for wavelet decomposition and statistical calculations.
*   **gopkg.in/yaml.v2:** YAML library for reading and writing configuration files, allowing for easy customization of the system.
*   **[Choose your database library here, e.g., "github.com/jmoiron/sqlx":** Used for storing tick data, signal data, and backtesting results. Provides a robust and efficient way to manage large datasets.]

## Installation

1.  **Prerequisites:** Ensure that Go (version 1.18 or higher) is installed and configured on your system. Also, make sure you have a working Go environment with `GOPATH` and `GOROOT` properly set.
2.  **Clone the Repository:**
    
3.  **Install Dependencies:**
    
4.  **Database Setup:** Create a database (e.g., PostgreSQL, MySQL) and configure the database connection parameters in the `config.yaml` file.
5.  **Build the Project:**
    

## Configuration

The system is configured using a `config.yaml` file. This file allows you to specify various parameters, including:

*   `database`: Database connection parameters (host, port, username, password, database name).
*   `tick_data`: Path to the tick data file (CSV format).
*   `wavelet`: Wavelet decomposition parameters (wavelet family, decomposition level).
*   `signal`: Signal generation parameters (thresholds for wavelet coefficient features).
*   `backtesting`: Backtesting parameters (initial capital, transaction costs, slippage).

Example `config.yaml`:

database:
  host: localhost
  port: 5432
  user: postgres
  password: password
  dbname: quantumforecaster

tick_data:
  path: data/EURUSD_tick_data.csv

wavelet:
  family: db4
  level: 5

signal:
  threshold_variance: 0.01

backtesting:
  initial_capital: 10000
  transaction_cost: 0.0001
  slippage: 0.00005

## Usage

To run the system, execute the compiled binary:

./quantumforecaster -config config.yaml

This will load the configuration from `config.yaml`, ingest the tick data, perform wavelet decomposition, generate trading signals, and run the backtesting engine.

Example code for ingesting tick data:

// sample code - not an actual code block, just indicative
package main

import (
 "fmt"
 "github.com/jjfhwang/QuantumForecaster/data"
)

func main() {
 // Example usage of the data ingestion module
 tickData, err := data.LoadTickData("data/EURUSD_tick_data.csv")
 if err != nil {
  fmt.Println("Error loading tick data:", err)
  return
 }

 fmt.Println("Number of ticks loaded:", len(tickData))
}

## Contributing

We welcome contributions to this project. Please follow these guidelines:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Implement your changes and write appropriate unit tests.
4.  Submit a pull request with a clear description of your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/jjfhwang/QuantumForecaster/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to thank the open-source community for providing the libraries and tools used in this project. Special thanks to the developers of `gonum.org/v1/gonum` for their excellent numerical computing library.