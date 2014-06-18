#include <iostream>
#include <fstream>
#include <sstream>
#include <array>

const int boardWidth = 256;
const int boardHeight = 256;

void displayMatrix(std::array<std::array<int, boardWidth>, boardHeight> &board) {
    /* Display filled array */
    for(int i = 0; i < boardHeight; ++i){
        for(int j = 0; j < boardWidth; ++j){
            std::cout<<board[i][j]<<" ";
        }
        std::cout<<std::endl;
    }
}

void setRow(std::array<std::array<int, boardWidth>, boardHeight> &board, int i, int x) {
    // Fill all cells on row i with value x
    for(int j = 0; j < boardWidth; ++j){
        board[i][j] = x;
    }
}

void setCol(std::array<std::array<int, boardWidth>, boardHeight> &board, int j, int x) {
    // Fill all cells on col j with value x
    for(int i = 0; i < boardHeight; ++i){
        board[i][j] = x;
    }
}

int queryRow(std::array<std::array<int, boardWidth>, boardHeight> &board, int i) {
    // Output the sum of all values on row i 
    int sum = 0;
    for(int j = 0; j < boardWidth; ++j){
        sum += board[i][j];
    }
    return sum;
}

int queryCol(std::array<std::array<int, boardWidth>, boardHeight> &board, int j) {
    // Output the sum of all values on column j 
    int sum = 0;
    for(int i = 0; i < boardHeight; ++i){
        sum += board[i][j];
    }
    return sum;
}

int main(int argc, char *argv[])
{
    std::ifstream file;
    std::string lineBuffer;


    std::array<std::array<int, boardWidth>, boardHeight> board;

    if (argc < 2) {
        std::cout << "Please add the name of the file to process." << std::endl;
        exit(1);
    }

    file.open(argv[1], std::ifstream::in);

    /* Row-Wise creation of array */
    for(int i = 0; i < boardHeight; ++i ){
        std::array<int, boardWidth> row;       //Row-array
        board[i] = row;
    }


    while (!file.eof()) 
    {
            getline(file, lineBuffer);
            if (lineBuffer.length() == 0)
                continue;
            else 
            {

                std::istringstream iss(lineBuffer);
                std::string cmd;
                int pos;
                int val = 0;
                iss >> cmd >> pos;
                if (cmd == "SetCol") {
                    iss >> val;
                    //std::cout << cmd << " " << pos << " " << val << std::endl;
                    setCol(board, pos, val);
                    //displayMatrix(board);
                } else if (cmd == "SetRow") {
                    iss >> val;
                    //std::cout << cmd << " " << pos << " " << val << std::endl;
                    setRow(board, pos, val);
                    //displayMatrix(board);
                } else if (cmd == "QueryCol") {
                    val = queryCol(board, pos);
                    //std::cout << cmd << " " << pos << " " << val << std::endl;
                    std::cout << val << std::endl;
                } else if (cmd == "QueryRow") {
                    val = queryRow(board, pos);
                    //std::cout << cmd << " " << pos << " " << val << std::endl;
                    std::cout << val << std::endl;
                } else {
                    std::cout << "Wrong!" << std::endl;
                }
            }
    }
    return 0;
}
