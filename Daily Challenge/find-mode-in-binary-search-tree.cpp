/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
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
