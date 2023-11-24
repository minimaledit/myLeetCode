class Solution {
public:
    int maxCoins(vector<int>& piles) {
        sort(piles.begin(), piles.end());
        int l = piles.size();
        int k = l / 3;
        int ans = 0;
        int c = piles.size() - 2;
        while(k != 0){
            ans += piles[c];
            c = c - 2;
            k--;
        }
        return ans;
    }
};
