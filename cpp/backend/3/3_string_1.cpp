#include <iostream>

class String{
    public:
        String(const char* str = NULL);
        String(const String& other);
        virtual ~String();

        String& operator= (const String& other);
        String& operator+ (const String& other);
        bool operator ==(const String& other);
        int getLength();

    private:
        char* m_date;
}