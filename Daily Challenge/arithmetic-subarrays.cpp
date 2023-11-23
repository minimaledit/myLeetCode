class Solution {
public:
    vector<bool> checkArithmeticSubarrays(vector<int>& nums, vector<int>& l, vector<int>& r) {
        vector<bool> ans;
        for(int i = 0; i < l.size(); ++i){
            vector<int> a;
            for(int j = l[i];j <= r[i];++j){
                a.push_back(nums[j]);
            }
            ans.push_back(func(a));
        }
        return ans;
    }
    bool func(vector<int>& arr){
        sort(arr.begin(),arr.end());
        int t = arr[1] - arr[0];
        for(int i = 2; i < arr.size(); ++i){
            if(arr[i] - arr[i - 1] != t){
                return false;
            }
        }
        return true;
    }
};
