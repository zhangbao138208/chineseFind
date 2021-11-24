<template>
    <div id="phoneCsPlusMinusReport">
        <!-- 电话客服加减分类报表 卡片区 -->
        <el-card class="box-card">
            <!-- 表单区 -->
            <el-form calss="form-query" inline ref="form" size="small" :model="queryMap" :rules="rules">
                <el-row class="query-row">
                    <el-form-item class="form-item" label="时间区间">
<!--                        <el-date-picker-->
<!--                                class="datePicker"-->
<!--                                :clearable="false"-->
<!--                                v-model="queryMap.times"-->
<!--                                type="datetimerange"-->
<!--                                value-format="yyyy-MM-dd HH:mm:ss"-->
<!--                                range-separator="至"-->
<!--                                start-placeholder="开始日期"-->
<!--                                end-placeholder="结束日期"-->
<!--                                :editable="false"-->
<!--                                ref="myDatePicker"-->
<!--                                :default-time="queryDefaultTime"-->
<!--                                @change="setDialogRange()"-->
<!--                        ></el-date-picker>-->
                      <el-date-picker
                          v-model="queryMap.times[0]"
                          type="date"
                          style="width: 130px"
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
                          style="width: 130px"
                          placeholder="结束时间">
                      </el-date-picker>
                    </el-form-item>
                    <el-form-item class="form-item" label="部门" prop="department"  label-width="40px">
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
                    <el-form-item class="form-item" label="客服名称" label-width="65px">
                        <el-select class="select"
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
                    <el-form-item label="客服类型" label-width="65px">
                        <el-select
                                @change="transferProductChange"
                                v-model="queryMap.csType"
                                placeholder="请选择"
                        >
                            <el-option :label="'普通客服'" :value="1"></el-option>
                            <el-option :label="'销售客服'" :value="12"></el-option>
                            <el-option :label="'VIP客服'" :value="13"></el-option>
                        </el-select>
                    </el-form-item>
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

            <el-table :cell-style="{padding:'0'}" :data="csData"  :header-cell-class-name="'table-header'" :row-class-name="tableRowClassName" border height="515px" size="small" style="width: 100%" v-loading="loading">
                <el-table-column prop="agentName" label="客服名" align="center" width="130">
                    <template slot-scope="scope">
                        {{null == scope.row.agentName?'-':scope.row.agentName}}
                    </template>
                </el-table-column>
                <el-table-column prop="openAccountCount" label="开户" align="center" width="50">
                    <template slot-scope="scope">
                        {{getValue(scope.row.openAccountCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="depositCount" label="存款" align="center" width="50">
                    <template slot-scope="scope">
                        {{getValue(scope.row.depositCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="upgradeCount" label="晋级" align="center" width="50">
                    <template slot-scope="scope">
                        {{getValue(scope.row.upgradeCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="stayLevelCount" label="挽留" align="center" width="50">
                    <template slot-scope="scope">
                        {{getValue(scope.row.stayLevelCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="cancellationCount" label="注销" align="center" width="50">
                    <template slot-scope="scope">
                        {{getValue(scope.row.cancellationCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="efficientCount" label="效率" align="center" width="50">
                    <template slot-scope="scope">
                        {{getValue(scope.row.efficientCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="salesMindsetCount" label="意识" align="center" width="50">
                    <template slot-scope="scope">
                        {{getValue(scope.row.salesMindsetCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="salesTechCount" label="技巧" align="center" width="50">
                    <template slot-scope="scope">
                        {{getValue(scope.row.salesTechCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="otherPlusCount" label="其他加分" align="center" width="70">
                    <template slot-scope="scope">
                        {{getValue(scope.row.otherPlusCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="phoneMannerCount" label="加分总和" align="center" width="70">
                    <template slot-scope="scope">
                        {{getValue(scope.row.plusCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="" label=" " align="center" width="8px" class-name="bg-gray">
                </el-table-column>
                <el-table-column prop="phoneMannerCount" label="电话礼仪" align="center" width="70">
                    <template slot-scope="scope">
                        {{getValue(scope.row.phoneMannerCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="expressionCount" label="语言表达" align="center" width="70">
                    <template slot-scope="scope">
                        {{getValue(scope.row.expressionCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="bussKnowledgeCount" label="业务知识" align="center" width="70">
                    <template slot-scope="scope">
                        {{getValue(scope.row.bussKnowledgeCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="bussFlowCount" label="业务流程" align="center" width="70">
                    <template slot-scope="scope">
                        {{getValue(scope.row.bussFlowCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="serviceEfficiencyCount" label="服务效率" align="center" width="70">
                    <template slot-scope="scope">
                        {{getValue(scope.row.serviceEfficiencyCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="serviceMindsetCount" label="营销意识" align="center" width="70">
                    <template slot-scope="scope">
                        {{getValue(scope.row.serviceMindsetCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="serviceTechniquesCount" label="营销技巧" align="center" width="70">
                    <template slot-scope="scope">
                        {{getValue(scope.row.serviceTechniquesCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="serviceAttitudeCount" label="服务态度" align="center" width="70">
                    <template slot-scope="scope">
                        {{getValue(scope.row.serviceAttitudeCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="otherMinusCount" label="其他扣分" align="center" width="70">
                    <template slot-scope="scope">
                        {{getValue(scope.row.otherMinusCount, 'Num')}}
                    </template>
                </el-table-column>
                <el-table-column prop="callTotalAverageSpan" label="扣分总和" align="center" width="70">
                    <template slot-scope="scope">
                        {{getValue(scope.row.minusCount, 'Num')}}
                    </template>
                </el-table-column>
            </el-table>
        </el-card>

    </div>
</template>

<script>
    import moment from 'moment';
    import { pickBy } from 'lodash';
    import axios from "axios";
    export default {
        data() {
            return {
                // 查询条件
                queryMap: {
                    times: [],
                    department: this.$store.state.userInfo.productIdArray[0],
                    csName: '',
                    csId:'',
                    csType:1
                },
                rules:{},
                departmentList: this.$store.state.userInfo.productIdArray, // 部门清单
                csList: [], // 客服清单
                dateLimit: 365, // 时间区间限制n天
                datedefault: 1, // 时间预设往前推n天
                loading: false,
                searchBtnLoading: false,
                // 资料-客服
                csData: [],
                csListLoading: false // 筛选条件-客服名称载入状态
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

                this.$set(self.queryMap.times, 0,endTime)
                this.$set(self.queryMap.times, 1, createTime)

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
            return true
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
                    csType:1
                }
                this.defaultDate()
                this.transferProductChange()
            },
            /**
             * 首次取得客服数据
             */
            async firstGetCsList(){
                this.loading = true;
                await this.transferProductChange()
                this.searchData()
            },
            async transferProductChange(){
                if(null== this.queryMap.department || '' === this.queryMap.department){
                    return;
                }
                this.csListLoading = true
                this.searchBtnLoading = true
                let params = {
                    department:this.queryMap.department,
                    startTime:moment(this.queryMap.times[0]).format('YYYY-MM-DD 00:00:00'),
                    endTime:moment(this.queryMap.times[1]).format('YYYY-MM-DD 23:59:59'),
                    csType:this.queryMap.csType
                }
                const { data: res } = await this.$http.get("system/agentServiceSound/getCs",{params});
                if(!res.success){
                    this.$message.error("获取在线客服数据失败:" + res.data.errorMsg);
                } else {
                    this.csList = res.data
                    this.queryMap.csId = null
                }
                this.csListLoading = false
                this.searchBtnLoading = false
            },
          downExcel() {
            if (!this.checkDialogRange()){
              return
            }
            let self = this;
            let params = {
              departments: self.queryMap.department
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
                list.push(item.id)
              })
              params.agentIds = this.csList.map(i=>i.rawId).toString()
            } else {
              params.agentIds = self.queryMap.csId.split(":")[0]
              params.agentName = self.queryMap.csId.split(":")[1]
            }
            params = pickBy(params, (item, key) => {
              return item !== '' && key !== 'times';
            });
            const $this = this;
            axios.request({
              url: "system/phoneCsReport/getPhoneCsPlusMinusReport/export",
              method: "get",
              responseType: "blob",
              params: params
            })
                .then(res => {
                  if (res.data.type === "application/json") {
                    return $this.$message.error(
                        "当前用户没有电话客服加减分类导出权限"
                    );
                  }
                  const data = res.data;
                  let url = window.URL.createObjectURL(data); // 将二进制文件转化为可访问的url
                  const a = document.createElement("a");
                  document.body.appendChild(a);
                  a.href = url;
                  a.download = "电话客服加减分类.xlsx";
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
                    departments: self.queryMap.department
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
                        list.push(item.id)
                    })
                    params.agentIds = this.csList.map(i=>i.rawId).toString()
                } else {
                    params.agentIds = self.queryMap.csId.split(":")[0]
                    params.agentName = self.queryMap.csId.split(":")[1]
                }
                params = pickBy(params, (item, key) => {
                    return item !== '' && key !== 'times';
                });
                let { data: res } = await self.$http.get(
                    'system/phoneCsReport/getPhoneCsPlusMinusReport', {params}
                );

                self.loading = false;
                self.searchBtnLoading = false

                if (res.success) {
                    self.csData = res.data;
                } else {
                    self.$message.error({
                        message: res.data.errorMsg,
                        type: 'error',
                    });
                }
            },
            getValue(item, type) {
                if(item && item !== "NaN" ) {
                    let num = parseFloat(Number(item).toFixed(2))
                    return type == 'Rate' ? num + '%' : num
                }
                return ''
            },
        },
        created() {
            this.defaultDate()
            this.firstGetCsList()
        }
    };
</script>
<style scoped>

    .el-select{
        width: 135px !important;
    }

    .datePicker{
        width: 335px;
    }

    .bg-gray {
        background: #f8f8f8;
    }
</style>
