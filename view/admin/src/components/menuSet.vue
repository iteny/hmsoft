<template>
  <div>
    <!-- 导航区域 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>活动管理</el-breadcrumb-item>
      <el-breadcrumb-item>活动列表</el-breadcrumb-item>
      <el-breadcrumb-item>活动详情</el-breadcrumb-item>
    </el-breadcrumb>
    <!-- 卡片区域 -->
    <el-card class="box-card">
      <!-- 通过基础的 1/24 分栏任意扩展组合形成较为复杂的混合布局。 -->
      <el-row :gutter="20">
        <el-col :span="16">
          <div class="grid-content bg-purple">
            <el-input placeholder="请输入内容"></el-input>
          </div>
        </el-col>
        <el-col :span="8">
          <div class="grid-content bg-purple">
            <el-button type="primary">主要按钮</el-button>
          </div>
        </el-col>
      </el-row>
      <el-table
        :data="menuList"
        style="width: 100%;margin-bottom: 20px;"
        row-key="id"
        border
        default-expand-all
        :tree-props="{children: 'children', hasChildren: 'hasChildren'}"
      >
        <el-table-column prop="name" label="菜单名称" sortable width="180"></el-table-column>
        <el-table-column prop="date" label="日期" sortable width="180"></el-table-column>
        <el-table-column prop="address" label="地址"></el-table-column>
      </el-table>
    </el-card>
  </div>
</template>
<script>
export default {
  data() {
    return {
      menuList: [],
    };
  },
  methods: {
    async getMenuList() {
      const res = await this.$http.post("/admin/menuList", {});
      this.menuList = res.data.data;
    },
  },
  created() {
    this.getMenuList();
  },
};
</script>
