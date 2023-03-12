#include <iostream>
#include <cstdlib>
#include <cstdio>
#include <vector>
#include <list>
#include <map>
#include <unordered_map>
#include <unordered_set>
#include <set>
#include <queue>
#include <deque>
#include <stack>
#include <ctime>
#include <random>
using namespace std;

int main() {
	time_t start = clock();
    cout << "hello, world!" << endl;
	srand(time(0));
	int num = 0;
	for (int i = 0; i < 1e6; i++) {
		num += num < 1000 ? rand() % 3 : -(rand() % 3);
	}
	cout << num << endl;
	time_t end = clock();
	cout << "运行时间：" << (end - start) / (CLOCKS_PER_SEC / 1000) << " ms" << endl;
    return 0;
}