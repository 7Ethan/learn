#include <iostream>

int main(int argc, char const *argv[])
{
    int arr[] = {34,56,234,65,344,56,345,4};
    int *p;
    int size = sizeof(arr);
    std::cout << size << std::endl;
    std::cout << "sizeof(&arr):" << sizeof(&arr) << std::endl;
    std::cout << "sizeof(p):" << sizeof(p) << std::endl;

    return 0;
}
