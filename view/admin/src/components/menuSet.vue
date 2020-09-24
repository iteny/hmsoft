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
            <!-- <el-button type="primary" @click="getElements">主要按钮</el-button> -->
          </div>
        </el-col>
      </el-row>
      <form class="sortme">
        <el-table
          ref="tableme"
          v-loading="loading"
          :data="menuList"
          style="width: 100%; margin-bottom: 20px"
          row-key="id"
          border
          default-expand-all
          :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
        >
          <el-table-column
            prop="name"
            label="菜单名称"
            sortable
            width="160"
          ></el-table-column>
          <el-table-column align="center" label="图标" width="80">
            <template slot-scope="scope">
              <i :class="scope.row.icon" style="font-size: 18px"></i>
            </template>
          </el-table-column>
          <el-table-column align="center" label="排序" width="80">
            <template slot-scope="scope">
              <!-- <el-input v-model="scope.row.sort" placeholder="请输入内容"></el-input> -->
              <input
                :name="scope.row.id"
                :value="scope.row.sort"
                style="
                  width: 30px;
                  text-align: center;
                  border-radius: 4px;
                  border-color: #66b1ff;
                "
              />
            </template>
          </el-table-column>
          <el-table-column prop="path" label="路径"></el-table-column>
          <el-table-column
            prop="isshow"
            label="是否显示"
            width="80"
            align="center"
          >
            <template slot-scope="scope">
              <el-switch
                v-model="scope.row.isshow"
                :active-value="1"
                :inactive-value="0"
                active-color="#13ce66"
                inactive-color="#ff4949"
                @change="changeSwitch($event, scope.row.id, scope.row.isshow)"
                >></el-switch
              >
            </template>
          </el-table-column>
          <el-table-column align="right">
            <template slot="header" slot-scope="scope">
              <el-button
                icon="el-icon-add-location"
                type="success"
                @click="addTopMenuDialog = true"
                >添加顶级菜单</el-button
              >
              <el-button
                icon="el-icon-add-location"
                type="success"
                @click="ajaxsort"
                >排序</el-button
              >
              <el-button
                icon="el-icon-loading"
                type="success"
                @click="reloadlocal"
                >刷新</el-button
              >
            </template>
            <template slot-scope="scope">
              <el-button
                icon="el-icon-edit"
                type="primary"
                @click="handleEdit(scope.$index, scope.row)"
                >编辑</el-button
              >
              <el-button
                icon="el-icon-delete"
                type="danger"
                @click="handleDelete(scope.$index, scope.row)"
                >删除</el-button
              >
            </template>
          </el-table-column>
        </el-table>
      </form>
    </el-card>
    <el-dialog
      title="添加顶级菜单"
      :visible.sync="addTopMenuDialog"
      width="50%"
    >
      <span>这是一段信息</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="addTopMenuDialog = false">取 消</el-button>
        <el-button type="primary" @click="addTopMenuDialog = false"
          >确 定</el-button
        >
      </span>
    </el-dialog>
  </div>
</template>
<script>
export default {
  data() {
    return {
      menuList: [],
      loading: true,
      sortList: [{}],
      addTopMenuDialog: false,
    };
  },
  inject: ["reload"],
  methods: {
    async getMenuList() {
      const res = await this.$http.post("/admin/menuList", {});
      this.loading = false;
      this.menuList = res.data.data;
    },
    async changeSwitch(ev, id, isshow) {
      if (ev == isshow) {
        const res = await this.$http.post("/admin/changeMenuIsshow", {
          id: id,
          isshow: ev,
        });
        if (res.data.status == "success") {
          this.$message.success("显示状态修改成功！");
        } else {
          this.$message.error("显示状态修改失败！");
        }
      }
    },
    async ajaxsort() {
      var form = $("form.sortme");
      var param = {};
      var sorts = form.serializeArray();
      $.each(sorts, function () {
        if (param[this.name] !== undefined) {
          if (!param[this.name].push) {
            param[this.name] = [param[this.name]];
          }
          param[this.name].push(this.value || "");
        } else {
          param[this.name] = this.value || "";
        }
      });
      var data = JSON.stringify(param);
      const res = await this.$http.post("/admin/sortMenu", { data: data });
      if (res.data.status == "success") {
        this.reload();
        this.$message.success("排序成功！");
      } else {
        this.$message.error("排序失败！");
      }
    },
    reloadlocal() {
      this.reload();
    },
  },
  created() {
    this.getMenuList();
  },
  // activated() {
  //   this.getMenuList();
  // },
};
</script>
