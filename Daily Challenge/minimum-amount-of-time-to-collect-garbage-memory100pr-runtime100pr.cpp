class Solution {
public:
    int garbageCollection(vector<string>& garbage, vector<int>& travel) {
         ios::sync_with_stdio(false);
        cin.tie(0);
        cout.tie(0);
        char m,p,g;
        int maxM = 0;
        int maxP = 0;
        int maxG = 0;
        int ans = 0;
        for(int i = 0;i < garbage.size();++i){
            m = 0;
            p = 0;
            g = 0;
            for(char c : garbage[i]){
                if(c == 'M'){
                    m++;
                }else{
                    if(c == 'P'){
                        p++;
                    }else if (c == 'G'){
                        g++;
                    }
                }
            }
            ans+= m + p + g;
            if(m!=0){
                maxM=i;
            }
            if(p!=0){
                maxP=i;
            }
            if(g!=0){
                maxG=i;
            } 
        }
        for(int i = 1;i <= maxM; ++i){
            ans+=travel[i-1];
        }
        for(int i = 1;i <= maxP; ++i){
            ans+=travel[i-1];
        }
        for(int i = 1;i <= maxG; ++i){
            ans+=travel[i-1];
        }
        return ans;
    }
};
