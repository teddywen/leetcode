package com.teddy.leetcode;

import com.teddy.leetcode.preset.TreeNode;

import java.util.Objects;

/**
 * @link https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/
 */
class Solution783 {

    public static void main(String[] args) {
        TreeNode root = new TreeNode(4);
        TreeNode l1 = new TreeNode(2);
        TreeNode r1 = new TreeNode(6);
        TreeNode l2 = new TreeNode(1);
        TreeNode r2 = new TreeNode(3);
        root.left = l1;
        root.right = r1;
        l1.left = l2;
        l1.right = r2;
        System.out.println(new Solution783().minDiffInBST(root));
    }

    private int last = -1;
    private int result = Integer.MAX_VALUE;

    public int minDiffInBST(TreeNode root) {
        dfs(root);
        return result;
    }

    public void dfs(TreeNode root) {
        if (Objects.isNull(root)) {
            return;
        }
        dfs(root.left);
        if (last > -1) {
            result = Math.min(root.val - last, result);
        }
        last = root.val;
        dfs(root.right);
    }
}