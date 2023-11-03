class Solution {
public:
    vector<string> buildArray(vector<int>& target, int n) {
        int index = 0;
        int targetSize = target.size();
        vector<string> ans;
        for(int currentNumber = 1; currentNumber <= n; currentNumber++){
            if(target[index] == currentNumber){
                ans.push_back("Push");
                index++;
            }else {
                ans.push_back("Push");
                ans.push_back("Pop");
            }
            if(index == targetSize){
                break;
            }
        }
        return ans;
    }
};
