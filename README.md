# solana_wallets_generator


**A simple script to generate solana wallets using either Python or Golang**

![Static Badge](https://img.shields.io/badge/Language-python-blue) ![Static Badge](https://img.shields.io/badge/Language-go-blue)

## Python Installation:

1. Clone the repository
```bash
git clone https://github.com/BlathanAevon/solana_wallets_generator.git
```

3. Go to directory
```bash
cd solana_wallets_generator/python
```
or (on Windows)
```bash
cd solana_wallets_generator\python
```
4. Virtual Environment creation
```bash
python (or python3) -m venv venv
```
5. Activate the virtual environment
    - Windows
      ```bash
      venv/Scripts/activate
      ```
    - Linux and MacOS
      ```bash
      source venv\bin\activate
      ```
6. Install dependencies (modules)
```bash
pip install -r requirements.txt
```

7. Set preferable amount of wallets in `main.py`
   ```python
   WALLETS_AMOUNT = <number of wallets>
   ```

## Run the program
```bash
python main.py
```


# Golang Installation:

1. Clone the repository
```bash
git clone https://github.com/BlathanAevon/solana_wallets_generator.git
```

3. Go to directory
     - Windows
      ```bash
      cd solana_wallets_generator\golang
      ```
    - Linux and MacOS
      ```bash
      cd solana_wallets_generator/golang
      ```
4. Install dependencies
```bash
go mod download
```

5. Set the amount of wallets in `gen.go`
   ```go
   const WALLETSAMOUNT int = <amount>
   ```

## RUN
```bash
go run .
```

