# GoRadio

GoRadio is a simple web application written in Golang for streaming radio stations. It also features a frontend with multiple radio stations and a horoscope feature that updates daily at 7 o'clock.

## Features

- Stream radio stations
- Multiple radio stations available
- Horoscope feature updated daily at 7 o'clock
- Access to application metrics on port 2112
- SQLite database integration

## Technologies Used

- Golang
- HTML/CSS
- JavaScript
- SQLite

## Usage

### Running the Application

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/goradio.git
    ```

2. Navigate to the project directory:

    ```bash
    cd goradio
    ```

3. Build and run the application:

    ```bash
    go build
    ./goradio
    ```

4. Open your web browser and visit `http://localhost:8080` to access the application.

### Accessing Metrics

- Application metrics can be accessed on port 2112. Open your web browser and visit `http://localhost:2112` to view metrics.

### Usage Notes

- Choose a radio station from the available options.
- Horoscope updates will occur automatically daily at 7 o'clock.

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -am 'Add your feature'`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Create a new Pull Request.

Please ensure that your pull request adheres to the code standards and includes appropriate documentation.
