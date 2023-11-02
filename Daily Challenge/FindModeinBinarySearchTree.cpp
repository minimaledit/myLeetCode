using namespace std;
class Solution {
public:
    map<int, int> counter;
    void dfs(TreeNode* node) {
        if (node == nullptr) {
            return;
        }
        counter[node->val]++;
        dfs(node->left);
        dfs(node->right);
    }
    vector<int> findMode(TreeNode* root) {
        dfs(root);
        int maxCount = 0;
        for (auto& pair : counter) {
            if (maxCount < pair.second) {
                maxCount = pair.second;
            }
        }
        vector<int> ans;
        for (auto& pair : counter) {
            if (maxCount == pair.second) {
                ans.push_back(pair.first);
            }
        }
        return ans;
    }
};