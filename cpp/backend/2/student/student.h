#ifndef STUDENT_H
#define STUDENT_H

class Student{
    public:
        Student();
        Student(int,int);
        ~Student();
        void display();
    private:
        int id;
        char name[20];
        int age;
};

#endif