class Solution {
public:
    int countNicePairs(vector<int>& nums) {
        ios_base::sync_with_stdio(false);
		cin.tie(0);
		cout.tie(0);
        map<int, long long> m;
        for(int i = 0;i < nums.size(); ++i){
            m[nums[i] - rev(nums[i])]++; 
        }
        long long ans = 0;
        for(auto t : m){
            if(t.second != 1){
                ans = (ans + t.second * (t.second - 1) / 2) % 1000000007;
            }
        }
        return ans;
    }
    int rev(int n){
        int res = 0;
        while(n > 0){
            res = res * 10 + n % 10;
            n /= 10;
        }
        return res;
    }
};
