# Load Testing Utility

This project aims to send GET requests to a simple HTTP server in order to retrieve the frequency of words listed in `search_words.csv` in a text file, downloaded from Gutenberg project.

The goal is to measure and compare the Latency and Throughput of the [HTTP Server](https://github.com/siddhant-vij/Word-Frequency-Server) under different loads by varying the number of server threads from the load testing utility.

<br>

## Table of Contents

1. [Load Testing Utility](#load-testing-utility)
1. [Running the Tests](#running-the-tests)
1. [Notes](#notes)
1. [Contributing](#contributing)
1. [License](#license)

<br>

## Load Testing Utility

- This utility reads words from `resources/search_words.csv` and generates HTTP GET requests for each word.
- It supports running multiple threads to simulate concurrent testing requests.
- The tool measures the latency (time taken for each request) and throughput (total requests completed in a time frame) for each run.

<br>

## Running the Tests
- Start the server application - clone the [server repository](https://github.com/siddhant-vij/Word-Frequency-Server) and proceed with the next steps to run the server.
- Run the load testing tool with a specified number of server & testing threads.
- The tool will display the average latency and throughput for each run in the terminal - saving the results in `resources/results.csv`.
- Create charts to further compare and analyze the results.
- Check the server performance with specialized tools like JMeter. Compare & contrast with this utility tool.

<br>

## Notes
- Ensure that this server is running before starting the load testing utility.
- The server and load testing utility are designed for educational purposes to demonstrate the effects of multithreading on performance.

<br>

## Contributing
Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.
1. **Fork the Project**
2. **Create your Feature Branch**: 
    ```bash
    git checkout -b feature/AmazingFeature
    ```
3. **Commit your Changes**: 
    ```bash
    git commit -m 'Add some AmazingFeature'
    ```
4. **Push to the Branch**: 
    ```bash
    git push origin feature/AmazingFeature
    ```
5. **Open a Pull Request**

<br>

## License

Distributed under the MIT License. See [`LICENSE`](https://github.com/siddhant-vij/Load-Testing-Utility/blob/main/LICENSE) for more information.