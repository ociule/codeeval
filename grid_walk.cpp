/**
    This is NOT a correct solution! The output from countClimbingStairs overflows if the input n is large enough.
    To solve this, we would need a bigint library. Otherwise, the solution is a correct memoized fibonacci, FWIW.
*/

#include <map>
#include <vector>
#include <set>
#include <iostream>

using namespace std;

std::map<int, unsigned long long int> cache;


int digitsSum(int n) {
    if (n < 0) {
        n = n * -1;
    }
    int sum = 0;
    int d, r;
    while (n > 0) {
        d = n / 10;
        r = n % 10;
        sum += r;
        n -= r;
        if (d > 0) {
            n /= 10;
        }
  }
  return sum;
}

const int LIMIT = 19;

bool accessible(int x, int y) {
    return digitsSum(x) + digitsSum(y) <= LIMIT;
}

void explore(vector<pair<int, int> > &queue, set<pair<int, int> > &visited, int x, int y) {
    pair<int, int> p = make_pair(x, y);
    if (accessible(x, y) && visited.find(p) == visited.end()) {
        visited.insert(p);
        queue.push_back(p);
    }
}

int main(int argc, char *argv[])
{
    vector<pair<int, int> > queue;
    set<pair<int, int> > visited;

    pair<int, int> start = make_pair(0, 0);

    queue.push_back(start);
    visited.insert(start);

    for (int i = 0; i < queue.size(); ++i) {
        int x = queue[i].first;
        int y = queue[i].second;

        explore(queue, visited, x - 1, y);
        explore(queue, visited, x + 1, y);
        explore(queue, visited, x, y - 1);
        explore(queue, visited, x, y + 1);
    }

    std::cout << queue.size() <<  std::endl;
    return 0;
}
