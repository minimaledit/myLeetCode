
class Solution {
public:
    int countPalindromicSubsequence(string s) {
        vector<int> first = vector(26,-1);
        vector<int> last = vector(26,-1);
        int len = s.size();
        int ans = 0;

        for(int i = 0; i < len; ++i){
            int t = s[i] - 'a';
            if(first[t] == -1){
                first[t] = i;
            } 
            last[t] = i;
        }
        
        for(int i = 0; i < 26; ++i){
            if(first[i] != -1 and first[i]!=last[i]){
                unordered_set<char> let;
                for(int j = first[i] + 1; j < last[i]; ++j){
                    let.insert(s[j]);
                }        
                ans += let.size();
            }
        }
        return ans;
    }
};
