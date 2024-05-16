#include <bits/stdc++.h>
using namespace std;

struct segtree {
    int size;
    vector<long long> sums;

    void init(int n) {
        size = 1;
        while (size < n) size *= 2;
        sums.assign(2*size, 0LL);
    }

    void set(int i, int v, int x, int lx, int rx) {
        
    }

    void set(int i, int v) {

    }

};

int main() {
    ios::sync_with_stdio(false);

    int n, m;
    cin >> n >> m;

    segtree st;
    st.init(n);
    for (int i = 0; i < n; i++) {
        int v;
        cin >> v;
        st.set(i, v);
    }
    
    return 0;
}