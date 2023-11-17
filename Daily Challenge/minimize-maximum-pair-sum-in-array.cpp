class Solution {
public:
    int minPairSum(vector<int>& nums) {
        ios::sync_with_stdio(false);
        cin.tie(0);
        cout.tie(0);
        sort(nums.begin(),nums.end());
        int len = nums.size();
        int ans = nums[0];
        for(int i=0;i<len/2;++i){
            ans = max(ans,nums[i]+nums[len-i-1]);
        }
        return ans;
    }
};
