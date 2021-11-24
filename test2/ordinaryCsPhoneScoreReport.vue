<template>
  <div id="ordinaryCsPhoneScoreReport">
    <!-- 加减量 卡片区 -->
    <el-card class="box-card">
      <!-- 表单区 -->
      <el-form calss="form-query" inline ref="form" size="small" :model="queryMap" :rules="rules">
        <el-row class="query-row">
          <el-form-item class="form-item" label="时间区间">
<!--            <el-date-picker-->
<!--              class="datePicker"-->
<!--              :clearable="false"-->
<!--              v-model="queryMap.times"-->
<!--              type="datetimerange"-->
<!--              value-format="yyyy-MM-dd HH:mm:ss"-->
<!--              range-separator="至"-->
<!--              start-placeholder="开始日期"-->
<!--              end-placeholder="结束日期"-->
<!--              :editable="false"-->
<!--              :default-time="queryDefaultTime"-->
<!--              ref="myDatePicker"-->
<!--              @change="setDialogRange()">-->
<!--              </el-date-picker>-->
            <el-date-picker
                v-model="queryMap.times[0]"
                type="date"
                @change="setDialogRange()"
                :clearable="false"
                :editable="false"
                placeholder="开始时间">
            </el-date-picker>
            <span >  &nbsp;&nbsp;&nbsp;~&nbsp;&nbsp;&nbsp;</span>
            <el-date-picker
                v-model="queryMap.times[1]"
                type="date"
                @change="setDialogRange()"
                :clearable="false"
                :editable="false"
                placeholder="结束时间">
            </el-date-picker>
          </el-form-item>
          <el-form-item class="form-item" label="部门" prop="department"  label-width="60px">
            <el-select
              filterable
              v-model="queryMap.department"
              placeholder="请选择"
              class="select"
              @change="transferProductChange"
            >
              <el-option
                v-for="(dept, i) in departmentList"
                :label="dept"
                :key="i"
                :value="dept"
              >
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item class="form-item" label="客服名称" label-width="80px">
            <el-select
              clearable
              filterable
              v-model="queryMap.csId"
              placeholder="请选择"
              v-loading="csListLoading"
            >
              <el-option
                  v-for="(item,index) in csList"
                  :label="`${item.username}(${item.rawId})`"
                  :key="index"
                  :value="`${item.rawId}:${item.username}`">
              </el-option>
          </el-select>
        </el-form-item>
        
          <!-- <el-form-item style="margin-left:10px">
            <el-button  @click="setWeekRange()" type="info">周数据</el-button>
          </el-form-item> -->
          <el-form-item style="margin-left: 10px">
            <el-button  @click="reset" icon="el-icon-refresh">重置</el-button>
            <el-button @click="downExcel" v-hasPermission="'user:export'"  icon="el-icon-download" size="small">导出</el-button>
            <el-button
              :loading="searchBtnLoading"
              type="primary"
              icon="el-icon-search"
              @click="searchData()"
            >
              查询
            </el-button>
          </el-form-item>
        </el-row>
      </el-form>
      <!-- 加减量 表格区 -->
      <el-table :cell-style="{padding:'0'}" :data="csData" :header-cell-class-name="'table-header'" :row-class-name="tableRowClassName" border height="505px" size="small" style="width: 100%" v-loading="loading">
        <el-table-column prop="csName" label="客服名" align="center" width="120">
          <template slot-scope="scope">
            {{scope.row.agentName?scope.row.agentName:scope.row.agentId}}
            
          </template>
        </el-table-column>
        <el-table-column prop="intimateCount" label="贴心嘉奖" align="center">
          <template slot-scope="scope">
            {{getValue(scope.row.intimateCount, 'Num')}}
          </template>
        </el-table-column>
        <el-table-column prop="professionalCount" label="专业嘉奖" align="center">
          <template slot-scope="scope">
            {{getValue(scope.row.professionalCount, 'Num')}}
          </template>
        </el-table-column>
        <el-table-column prop="wiseCount" label="智慧嘉奖" align="center">
          <template slot-scope="scope">
            {{getValue(scope.row.wiseCount, 'Num')}}
          </template>
        </el-table-column>
        <el-table-column prop="greatCount" label="口头表扬" align="center">
          <template slot-scope="scope">
            {{getValue(scope.row.greatCount, 'Num')}}
          </template>
        </el-table-column>
        <el-table-column prop="hintCount" label="温馨提示" align="center">
          <template slot-scope="scope">
            {{getValue(scope.row.hintCount, 'Num')}}
          </template>
        </el-table-column>
        <el-table-column prop="warnCount" label="警告" align="center">
          <template slot-scope="scope">
            {{getValue(scope.row.warnCount, 'Num')}}
          </template>
        </el-table-column>
        <el-table-column prop="faultCount" label="失误" align="center">
          <template slot-scope="scope">
            {{getValue(scope.row.faultCount, 'Num')}}
          </template>
        </el-table-column>
        <el-table-column prop="accidentCount" label="事故" align="center">
          <template slot-scope="scope">
            {{getValue(scope.row.accidentCount, 'Num')}}
          </template>
        </el-table-column>
        <el-table-column prop="plusCount" label="加分量" align="center">
          <template slot-scope="scope">
            {{getValue(scope.row.plusCount, 'NumWithZero')}}
          </template>
        </el-table-column>
        <el-table-column prop="minusCount" label="错误量" align="center">
          <template slot-scope="scope">
            {{getValue(scope.row.minusCount, 'NumWithZero')}}
          </template>
        </el-table-column>
        <el-table-column prop="standardScore" label="基准分" align="center">
          <template slot-scope="scope">
            {{getValue(scope.row.standardScore, 'NumWithZero')}}
          </template>
        </el-table-column>
        <el-table-column prop="plusAndMinusScoreTotal" label="加减分得分" align="center">
          <template slot-scope="scope">
            {{getValue(scope.row.plusAndMinusScoreTotal, 'NumWithZero')}}
          </template>
        </el-table-column>
        <el-table-column prop="scoreTotal" label="总得分" align="center">
          <template slot-scope="scope">
            {{getValue(scope.row.scoreTotal, 'NumWithZero')}}
          </template>
        </el-table-column>
        <el-table-column prop="scoreTotal" label="备注" align="center">
          <template slot-scope="scope">
            <a href="javascript:void(0)"
            v-if="scope.row.agentId"
            @click="openRemarkDialog(scope.row)"
            >{{scope.row.remark?scope.row.remark:"无"}}</a>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
 
    <el-dialog :title="'备注'" :visible.sync="remarkDialog" width="35%" @close="closeDialog">
      <span>
        <el-form :model="remarkObj"
                 ref="addFormRef"
                 label-width="150px">
          <el-row>
            <el-col :span="18">
              <el-form-item label="备注" prop="holidayName">
                <el-input type="textarea" v-model="remarkObj.remark" rows="7" size="medium" maxlength="30"></el-input>
              </el-form-item>
            </el-col>
            
          </el-row>
          <el-row>
            <el-col :span="17">&nbsp;</el-col>
            <el-col :span="5">
              {{ remarkObj.remark?remarkObj.remark.length: 0 }}  / 30
            </el-col>
          </el-row>
        </el-form>
      </span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="remarkDialog = false">取 消</el-button>
        <el-button
            type="primary" 
            @click="saveOrUpdateRemark"
            :loading="btnLoading"
            :disabled="btnDisabled">
            确 定
        </el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import moment from 'moment';
import { pickBy, cloneDeep } from 'lodash';
import axios from "axios";

  export default {
    data() {
      return {
        // 查询条件
        queryMap: {
          times: [],
          department: this.$store.state.userInfo.productIdArray[0],
          csType:null,
          agentIds:[],
          csId:''
        },
        ckeckQueryMap: {
          department: '',
          csType:null,
          agentIds:[],
          csId:''
        },
        // 栏位验证规则
        rules: {}, 
        departmentList: this.$store.state.userInfo.productIdArray, // 部门清单
        csList: [], // 客服清单
        dateLimit: 365, // 时间区间限制n天
        datedefault: 1, // 时间预设往前推n天
        loading: false,
        searchBtnLoading: false,
        // 资料-客服
        csData: [],
        csListLoading: false, // 筛选条件-客服名称载入状态
        remarkDialog:false,
        btnLoading:false,
        btnDisabled:false,
        remarkObj:{
          id:null,
          csId:'',
          remark:null,
          reportType:null
        }
      }
    },
    methods: {
      /**
       * 最底列class
       */
      tableRowClassName({row, rowIndex}) {
        if(rowIndex == this.csData.length - 1) {
          return 'last-row'
        }
      },
      /**
       * 筛选-时间区间
       */
      setDialogRange() {
        let self = this;
        let createTime = self.queryMap.times[0];
        let endTime = self.queryMap.times[1];
        if (moment(createTime) >= moment(endTime)){
          this.$set(self.queryMap.times, 0, endTime)
          this.$set(self.queryMap.times, 1,createTime)
        }
        return
        let totoalDays = moment(endTime).diff(moment(createTime), 'days');
        if (totoalDays > self.dateLimit) {
          this.defaultDate()
          self.$alert('开始时间与结束时间间隔不能超过' + self.dateLimit + '天');
        }
        this.transferProductChange()
      },
      checkDialogRange() {
        let self = this;
        let createTime = self.queryMap.times[0];
        let endTime = self.queryMap.times[1];

        let totoalDays = moment(endTime).diff(moment(createTime), 'days');
        if (totoalDays > self.dateLimit) {
          this.defaultDate()
          self.$alert('开始时间与结束时间间隔不能超过' + self.dateLimit + '天');
          return false
        }
        this.transferProductChange()
        return  true
      },
      /**
       * 预设时间(当前时间往前推n日)
       */
      defaultDate() {
        let startDay= new Date()
        startDay.setDate(startDay.getDate() - this.datedefault)
        this.$set(this.queryMap.times,0,new Date(startDay.getFullYear(), startDay.getMonth(), 1, 0, 0, 0))
        this.$set(this.queryMap.times,1, new Date(startDay.getFullYear(), startDay.getMonth(), startDay.getDate(), 23, 59, 59))
      },
      /**
       * 重置
       */
      reset() {
        this.queryMap = {
          times: [],
          department: this.$store.state.userInfo.productIdArray[0],
          csId: '',
        }
        this.defaultDate()
        this.transferProductChange()
      },
      /**
       * transfer窗口的部门下拉列表改变事件
       */
      async transferProductChange(){
          if(null== this.queryMap.department || '' === this.queryMap.department){
              return;
          }
          this.csListLoading = true
          this.searchBtnLoading = true
          this.queryMap.csName = ''
          let params = {
            department:this.queryMap.department,
            startTime:moment(this.queryMap.times[0]).format('YYYY-MM-DD HH:mm:ss'),
            endTime:moment(this.queryMap.times[1]).format('YYYY-MM-DD HH:mm:ss'),
            csType:1
          }
          const { data: res } = await this.$http.get("system/agentServiceSound/getCs",{params});
          if(!res.success){
              this.$message.error("获取客服数据失败:" + res.data.errorMsg);
          } else {
            this.csList = cloneDeep(res.data)
            this.queryMap.csId = ''
          }
          this.csListLoading = false
          this.searchBtnLoading = false
      },
      /**
       * 首次取得客服数据
       */
      async firstGetCsList(){
          this.loading = true;
          await this.transferProductChange()
          this.searchData()
      },
      downExcel() {
        if (!this.checkDialogRange()){
          return
        }
        let self = this;

        let params = {
          department: self.queryMap.department
        }
        if (self.queryMap.times) {
          const start = moment(self.queryMap.times[0]).format('YYYY-MM-DD 00:00:00')
          const end = moment(self.queryMap.times[1]).format('YYYY-MM-DD 23:59:59')
          params.startTime = start || '';
          params.endTime = end || '';
        }
        if(!self.queryMap.csId || self.queryMap.csId == '') {
          let list = []
          self.csList.forEach(item => {
            list.push(item.rawId)
          })
          params.agentIds = list.toString()
        } else {
          const comb = self.queryMap.csId.split(':')
          params.agentIds = comb[0]
          params.agentName = comb[1]
        }
        params = pickBy(params, (item, key) => {
          return item !== '' && key !== 'times';
        });
        const $this = this;
        axios.request({
          url: "system/phoneCsReport/getOrdinaryCsPhoneScoreReport/export",
          method: "get",
          responseType: "blob",
          params: params
        })
            .then(res => {
              if (res.data.type === "application/json") {
                return $this.$message.error(
                    "当前用户没有普通客服电话得分导出权限"
                );
              }
              const data = res.data;
              let url = window.URL.createObjectURL(data); // 将二进制文件转化为可访问的url
              const a = document.createElement("a");
              document.body.appendChild(a);
              a.href = url;
              a.download = "普通客服电话得分.xlsx";
              a.click();
              document.body.removeChild(a)
            });
      },
      /**
       * 搜寻
       */
      async searchData() {
        if (!this.checkDialogRange()){
          return
        }
        let self = this;
        self.loading = true;
        self.searchBtnLoading = true
        let params = {
          department: self.queryMap.department
        }
        if (self.queryMap.times) {
          const start = moment(self.queryMap.times[0]).format('YYYY-MM-DD 00:00:00')
          const end = moment(self.queryMap.times[1]).format('YYYY-MM-DD 23:59:59')
          params.startTime = start || '';
          params.endTime = end || '';
        }
        if(!self.queryMap.csId || self.queryMap.csId == '') {
          let list = []
          self.csList.forEach(item => {
            list.push(item.rawId)
          })
          params.agentIds = list.toString()
        } else {
          const comb = self.queryMap.csId.split(':')
          params.agentIds = comb[0]
          params.agentName = comb[1]
        }
        params = pickBy(params, (item, key) => {
          return item !== '' && key !== 'times';
        });
        let { data: res } = await self.$http.get(
          'system/phoneCsReport/getOrdinaryCsPhoneScoreReport', {params}
        );

        self.loading = false;
        self.searchBtnLoading = false

        if (res.success) {
          self.ckeckQueryMap = cloneDeep(self.queryMap)
          self.csData = res.data;
        } else {
          self.$message.error({
            message: res.data.errorMsg,
            type: 'error',
          });
        }
      },
      getValue(item, type) {
        if((item && item !== "NaN")) {
          let num = parseFloat(Number(item).toFixed(2))
          return type == 'Rate' ? num + '%' : num
        } else if(item == 0 && type == 'NumWithZero'){
          return item
        }
        return ''
      },
      setWeekRange(){
        let date = new Date()
        let start, end;
        if(date.getDay() < 6){
            date.setDate(date.getDate() - date.getDay() - 2)
            end = new Date(date.getFullYear(), date.getMonth(), date.getDate(), 23, 59, 59)
        } else if(date.getDay() === 6){
            date.setDate(date.getDate() - 1)
            end = new Date(date.getFullYear(), date.getMonth(), date.getDate(), 23, 59, 59)
        } else if(date.getDay() === 7){
            date.setDate(date.getDate() - 2)
            end = new Date(date.getFullYear(), date.getMonth(), date.getDate(), 23, 59, 59)
        }
        date.setDate(date.getDate() - 6)
        start = new Date(date.getFullYear(), date.getMonth(), date.getDate(), 0, 0, 0)
        this.queryMap.times = [start,end]
        this.searchData()
      },
      openRemarkDialog(row){
        this.remarkObj.id = row.remarkId
        this.remarkObj.csId = row.agentId
        this.remarkObj.reportType = 1
        this.remarkObj.remark = row.remark
        this.remarkDialog = true
      },
      closeDialog(){},
      async saveOrUpdateRemark(){
        let { data: res } = await this.$http.post(
          'system/reportRemark', this.remarkObj
        );
        if(res.success){
          this.$notify.success({
              title:'操作成功',
          });
          this.searchData()
          this.remarkDialog = false
        } else{
          this.$message.error(res.data.data.errorMsg);
        }
      }
    },
    created() {
      this.defaultDate()
      this.firstGetCsList()
    }
  };
</script>

<style scoped>
.sub-title {
  font-size: 14px;
  margin-top: 30px;
  margin-bottom: 10px;
  font-weight: bold;
}
.el-table {
  margin-bottom: 16px;
}
.block-fliter {
  margin-bottom: 20px;
}
.el-checkbox {
  margin-right: 10px;
}
.el-select,
.el-input {
  width: 128px;
}
.datePicker {
  width: 335px;
}
.el-table >>> .last-row {
  background: #f8f8f8;
}
</style>

