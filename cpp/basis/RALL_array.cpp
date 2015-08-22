/**
 * RAII的做法是使用一个对象，在其构造时获取对应的资源，
 * 在对象生命期内控制对资源的访问，使之始终保持有效，
 * 最后在对象析构的时候，释放构造时获取的资源。
*/

/**
 * 具体表现（实现）就是：把类中所有分配了的资源，在析构函数中释放掉，就不再需要一一地手动释放。
*/

#include <iostream> 

using namespace std; 

class ArrayOperation { 
public : 
    ArrayOperation() { 
        m_Array = new int [10]; 
    } 
    void InitArray() { 
        for (int i = 0; i < 10; ++i) { 
            *(m_Array + i) = i; 
        } 
    } 

    void ShowArray() { 
        for (int i = 0; i <10; ++i) { 
            cout<<m_Array[i]<<endl; 
        } 
    } 

    ~ArrayOperation() { 
        cout<< "~ArrayOperation is called" <<endl; 
        if (m_Array != NULL ) { 
            delete[] m_Array; 
            m_Array = NULL ; 
        } 
    } 
private : 
    int *m_Array; 
};

bool OperationA();
bool OperationB(); 

int 
main() { 
    ArrayOperation arrayOp; 
    arrayOp.InitArray(); 
    arrayOp.ShowArray(); 
    return 0; 
}
