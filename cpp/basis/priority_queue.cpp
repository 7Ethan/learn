/*
 *本程序测试std::priority_queue的基本用法
 * */

#include <iostream>
#include <functional>
#include <queue>
#include <vector>

using namespace std;

struct Node{
  int data;
  bool operator<(const Node& node) const{
    return data < node.data;
  }
};

class cmp{
public:
  bool operator()(const Node& n1,const Node& n2) const{
    return n1 < n2;
  }
};

int main(){
  priority_queue<Node,vector<Node>,cmp> q;
  Node n1;n1.data = 5;
  Node n2;n2.data = 4;
  Node n3;n3.data = 3;
  Node n4;n4.data = 2;
  Node n5;n5.data = 1;

  q.push(n1);
  q.push(n2);
  q.push(n3);
  q.push(n4);
  q.push(n5);

  while(!q.empty()){
    cout << q.top().data << " ";
    q.pop();
  }
  cout << endl;
  return 0;
}

