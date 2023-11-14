class Solution {
public:
    bool isVowel(char c){
        if(c == 'a' || c == 'A' || c == 'e'|| c == 'E'|| c == 'o' 
            || c == 'O' || c == 'u' || c == 'U'|| c == 'i'|| c == 'I'){
                return true;
            }else return false;
    }
    string sortVowels(string s) {
        string vowels;
        for(char t : s){
            if(isVowel(t)){
                vowels+=t;
            }
        }
        sort(vowels.begin(),vowels.end());

        string ans;
        int counter=0;
        for(char t : s){
            if(isVowel(t)){
                ans+=vowels[counter];
                counter++;
            }else{
                ans+=t;
            }
        }
        return ans;
    }
};
