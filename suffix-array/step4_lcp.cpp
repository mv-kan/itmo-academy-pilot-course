#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

void count_sort(vector<int>&p, vector<int>&c) {
    int n = p.size();
    vector<int> count(n);
    for (auto x : c) {
        count[x]++;
    }
    
    vector<int> p_new(n);
    vector<int> pos(n);
    pos[0] = 0;
    for (int i = 1; i < n; i++) {
		pos[i] = pos[i-1] + count[i-1];
	}
    for (auto x : p) {
        int i = c[x];
        p_new[pos[i]] = x;
        pos[i]++;
    }
    p = p_new;
}
int main() {
    string s{};
    cin >> s;
    s += "$";
    int n = s.size();
    vector<int> p(n);
    vector<int> c(n);

    {
        vector< pair<char, int> > a(n);
        for (int i = 0; i < n; i++) {
            a[i] = pair{s[i], i};
        }
        sort(a.begin(), a.end());
        for (int i = 0; i < n; i++) {
            p[i] = a[i].second;
        }
        c[p[0]] = 0;
        for (int i = 1; i < n; i++) {
            if (a[i-1].first == a[i].first) {
                c[p[i]] = c[p[i-1]];
            } else {
                c[p[i]] = c[p[i-1]] + 1;
            }
        }
    }
    int k = 0;
    while (1 << k < n) {
        for (int i = 0; i < n; i++) {
			p[i] = (p[i] - (1<<k) + n) % n;
		}

        count_sort(p, c);

        vector<int> c_new(n);
        c_new[p[0]] = 0;
        for (int i = 1; i < n; i++) {
            pair<int, int> prev{c[p[i-1]], c[(p[i-1]+(1<<k))%n]};
            pair<int, int> now{c[p[i]], c[(p[i]+(1<<k))%n]};
            if (now == prev) {
                c_new[p[i]] = c_new[p[i-1]];
            } else {
                c_new[p[i]] = c_new[p[i-1]] + 1;
            }
        }
        c = c_new;
        k++;
    }
    vector<int> lcp(n);
    k = 0;
    for (int i = 0; i < n-1; i++) {
        int pi = c[i];
        int j = p[pi-1];
        while(s[i+k] == s[j+k]) k++;
        lcp[pi] = k;
        k = max(k-1, 0);
    }    
    // for (int i = 0; i < n; i++) {
    //     cout << p[i] << "\t"<<s.substr(p[i], s.size()) << "\n";
    // }
    for (int i = 0; i < n; i++) {
        cout << p[i] << " ";
    }
    cout << endl;
    for (int i = 1; i < n; i++) {
        cout << lcp[i] << " ";
    }
    cout << endl;
}