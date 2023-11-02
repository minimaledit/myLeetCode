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
class Solution {
public:
    int counter = 0;
    int averageOfSubtree(TreeNode* root) {
        finder(root);
        return counter;
    }

    pair<int, int> finder(TreeNode* root){
        if(root == nullptr){
            return {0,0};
        }

        pair<int, int> left = finder(root->left);
        pair<int, int> right = finder(root->right);
        
        int sum = left.first + right.first + root->val;
        int cnt = left.second + right.second + 1;
        //int avg =(left.first + right.first + root->val)/(left.second + right.second + 1);
        
        if(sum / cnt == root->val){
            counter++;
        }

        return {sum, cnt};
    }

};
