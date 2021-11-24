<template>
  <div id="phoneCsOverview">
    <!-- 综合报表卡片区 -->
    <el-card class="box-card">
      <!-- 表单区 -->
      <el-form inline ref="form" size="small" :model="queryMap" :rules="rules">
        <el-row class="query-row">
          <el-form-item label="选择部门">
            <el-checkbox-group v-model="queryMap.departments" size="small" :min="1">
              <el-checkbox v-for="item in departmentList" :label="item" :value="item" :key="item" border></el-checkbox>
            </el-checkbox-group>
          </el-form-item>
        </el-row>
        <el-row class="query-row">
          <el-form-item label="时间区间">
            <!--            <el-date-picker-->
            <!--              v-show="dateScale == 1"-->
            <!--              class="datePicker"-->
            <!--              :clearable="false"-->
            <!--              v-model="queryMap.times"-->
            <!--              :type="'daterange'"-->
            <!--              value-format="yyyy-MM-dd HH:mm:ss"-->
            <!--              range-separator="至"-->
            <!--              start-placeholder="开始日期"-->
            <!--              end-placeholder="结束日期"-->
            <!--              :default-time="queryDefaultTime"-->
            <!--              :editable="false"-->
            <!--              ref="myDatePicker"-->
            <!--              @change="setDialogRange()"-->
            <!--            ></el-date-picker>-->
            <el-date-picker
                v-show="dateScale == 1"
                v-model="queryMap.times[0]"
                type="date"
                @change="setDialogRange()"
                :clearable="false"
                :editable="false"
                placeholder="开始时间">
            </el-date-picker>
            <span  v-show="dateScale == 1">  &nbsp;&nbsp;&nbsp;~&nbsp;&nbsp;&nbsp;</span>
            <el-date-picker
                v-show="dateScale == 1"
                v-model="queryMap.times[1]"
                type="date"
                @change="setDialogRange()"
                :clearable="false"
                :editable="false"
                placeholder="结束时间">
            </el-date-picker>
            <!--            <el-date-picker-->
            <!--              v-show="dateScale != 1"-->
            <!--              class="datePicker"-->
            <!--              :clearable="false"-->
            <!--              v-model="queryMap.times"-->
            <!--              :type="'monthrange'"-->
            <!--              value-format="yyyy-MM-dd HH:mm:ss"-->
            <!--              range-separator="至"-->
            <!--              start-placeholder="开始日期"-->
            <!--              end-placeholder="结束日期"-->
            <!--              :default-time="queryDefaultTime"-->
            <!--              :editable="false"-->
            <!--              ref="myMonthPicker"-->
            <!--              @change="setDialogRange()"-->
            <!--            ></el-date-picker>-->
            <el-date-picker
                v-show="dateScale != 1"
                v-model="queryMap.times[0]"
                type="month"
                @change="setDialogRange()"
                :clearable="false"
                :editable="false"
                placeholder="开始时间">
            </el-date-picker>
            <span  v-show="dateScale != 1">  &nbsp;&nbsp;&nbsp;~&nbsp;&nbsp;&nbsp;</span>
            <el-date-picker
                v-show="dateScale != 1"
                v-model="queryMap.times[1]"
                type="month"
                @change="setDialogRange()"
                :clearable="false"
                :editable="false"
                placeholder="结束时间">
            </el-date-picker>
          </el-form-item>
          <!-- <el-form-item style="margin-left:10px" class="form-item" label="部门" prop="department">
            <el-select
              filterable
              clearable
              v-model="queryMap.department"
              @change="transferProductChange"
              placeholder="请选择"
              class="select">
              <el-option
                v-for="(dept, i) in departmentList"
                :label="dept"
                :key="i"
                :value="dept">
              </el-option>
            </el-select>
          </el-form-item> -->
          <el-form-item label="客服类型" label-width="70px">
            <el-select v-model="queryMap.csType" @change="transferProductChange" placeholder="请选择">
              <el-option :label="'普通客服'" :value="1"></el-option>
              <el-option :label="'销售客服'" :value="12"></el-option>
              <el-option :label="'VIP客服'" :value="13"></el-option>
            </el-select>
          </el-form-item>
         <el-form-item style="margin-left:10px">
           <el-button  @click="setDateScale(1)" :type="dateScale==1?'primary':''">按日</el-button>
           <el-button  @click="setDateScale(2)" :type="dateScale==2?'primary':''">按月</el-button>
         </el-form-item>
          <el-form-item style="margin-left:10px">
            <el-button @click="reset" icon="el-icon-refresh">重置</el-button>
            <el-button :loading="btnLoading" type="primary" icon="el-icon-search" @click="searchData()">
              查询
            </el-button>
          </el-form-item>
        </el-row>
      </el-form>
      <!-- 图表区 -->
      <!-- 图表-有效电话量-->
      <el-col :span="24">
        <div class="sub-title">
          <span>有效电话量</span>
        </div>
        <echart :id="'effectiveTalks'" :options.sync="effectiveTalkNumberData" v-loading="loading"></echart>
        <chartTable :resultDepartments="resultDepartments" :data="effectiveTalkNumberTable"></chartTable>
      </el-col>
      <!-- 图表-人均有效电话量-->
      <el-col :span="24">
        <div class="sub-title">
          <span>人均有效电话量</span>
        </div>
        <echart
          :id="'averageEffectiveTalkNumber'"
          :options.sync="averageEffectiveTalkNumberData"
          v-loading="loading"
        ></echart>
        <chartTable :resultDepartments="resultDepartments" :data="averageEffectiveTalkNumberTable"></chartTable>
      </el-col>
      <!-- 图表-人均月通话时长（时）-->
      <el-col :span="24">
        <div class="sub-title">
          <span>人均月通话时长（时）</span>
        </div>
        <echart
          :id="'averageCostTimePerPerson'"
          :options.sync="averageCostTimePerPersonData"
          v-loading="loading"
        ></echart>
        <chartTable :resultDepartments="resultDepartments" :data="averageCostTimePerPersonTable"></chartTable>
      </el-col>
      <!-- 图表-平均通话时长（秒）-->
      <el-col :span="24">
        <div class="sub-title">
          <span>平均通话时长（秒）</span>
        </div>
        <echart :id="'averageCostTime'" :options.sync="averageCostTimeData" v-loading="loading"></echart>
        <chartTable :resultDepartments="resultDepartments" :data="averageCostTimeTable"></chartTable>
      </el-col>
      <!-- 图表-电话呼入接通率-->
      <el-col :span="24">
        <div class="sub-title">
          <span>电话呼入接通率</span>
        </div>
        <echart :id="'validCallInput'" :options.sync="validCallInputData" v-loading="loading"></echart>
        <chartTable :resultDepartments="resultDepartments" :data="validCallInputTable" :rate="true"></chartTable>
      </el-col>
      <!-- 图表-电话呼出接通率-->
      <el-col :span="24">
        <div class="sub-title">
          <span>电话呼出接通率</span>
        </div>
        <echart :id="'validCallOutput'" :options.sync="validCallOutputData" v-loading="loading"></echart>
        <chartTable :resultDepartments="resultDepartments" :data="validCallOutputTable" :rate="true"></chartTable>
      </el-col>
      <!-- 图表-电话总接通率-->
      <el-col :span="24">
        <div class="sub-title">
          <span>电话总接通率</span>
        </div>
        <echart :id="'validCallTotal'" :options.sync="validCallTotalData" v-loading="loading"></echart>
        <chartTable :resultDepartments="resultDepartments" :data="validCallTotalTable" :rate="true"></chartTable>
      </el-col>
      <!-- 图表-电话嘉奖率-->
      <el-col :span="24">
        <div class="sub-title">
          <span>电话嘉奖率</span>
        </div>
        <echart :id="'plusRate'" :options.sync="plusRateData" v-loading="loading"></echart>
        <chartTable :resultDepartments="resultDepartments" :data="plusRateTable" :rate="true"></chartTable>
      </el-col>
      <!-- 图表-电话错误率-->
      <el-col :span="24">
        <div class="sub-title">
          <span>电话错误率</span>
        </div>
        <echart :id="'minusRate'" :options.sync="minusRateData" v-loading="loading"></echart>
        <chartTable :resultDepartments="resultDepartments" :data="minusRateTable" :rate="true"></chartTable>
      </el-col>
      <!-- 图表-电话营销意识错误次数-->
      <el-col :span="24">
        <div class="sub-title">
          <span>电话营销意识错误次数</span>
        </div>
        <echart :id="'serviceMindsetCount'" :options.sync="serviceMindsetCountData" v-loading="loading"></echart>
        <chartTable :resultDepartments="resultDepartments" :data="serviceMindsetCountTable"></chartTable>
      </el-col>
      <!-- 图表-电话营销技巧错误次数-->
      <el-col :span="24">
        <div class="sub-title">
          <span>电话营销技巧错误次数</span>
        </div>
        <echart :id="'serviceTechniqueCount'" :options.sync="serviceTechniqueCountData" v-loading="loading"></echart>
        <chartTable :resultDepartments="resultDepartments" :data="serviceTechniqueCountTable"></chartTable>
      </el-col>
      <!-- 图表-电话服务态度错误次数-->
      <el-col :span="24">
        <div class="sub-title">
          <span>电话服务态度错误次数</span>
        </div>
        <echart :id="'serviceAttitudeCount'" :options.sync="serviceAttitudeCountData" v-loading="loading"></echart>
        <chartTable :resultDepartments="resultDepartments" :data="serviceAttitudeCountTable"></chartTable>
      </el-col>
    </el-card>
  </div>
</template>

<script>
import chartTable from "@/components/chartTable";
import echart from "@/components/echart";
import { pickBy, cloneDeep } from "lodash";
import moment from "moment";
import constants from '@/utils/constants';
export default {
  data() {
    return {
      // 查询条件
      queryMap: {
        times: [],
        departments: [],
        department: null,
        csType: 1
      },
      resultDepartments: [],
      // 栏位验证规则
      rules: {},
      departmentList: [], // 部门清单
      dateLimit: 365, // 时间区间限制n天
      datedefault: 1, // 时间预设往前推n天
      loading: false,
      effectiveTalkNumberData: {}, // 有效电话量-资料
      effectiveTalkNumberTable: [], // 有效电话量-表格
      averageEffectiveTalkNumberData: {}, // 人均有效电话量-资料
      averageEffectiveTalkNumberTable: [], // 人均有效电话量-表格
      averageCostTimePerPersonData: {}, // 人均月通话时长（时）-资料
      averageCostTimePerPersonTable: [], // 人均月通话时长（时）-表格
      averageCostTimeData: {}, // 平均通话时长（秒）-资料
      averageCostTimeTable: [], // 平均通话时长（秒）-表格
      validCallInputData: {}, // 电话呼入接通率-资料
      validCallInputTable: [], // 电话呼入接通率-表格
      validCallOutputData: {}, // 电话呼出接通率-资料
      validCallOutputTable: [], // 电话呼出接通率-表格
      validCallTotalData: {}, // 电话总接通率-资料
      validCallTotalTable: [], // 电话总接通率-表格
      plusRateData: {}, // 电话嘉奖率-资料
      plusRateTable: [], // 电话嘉奖率-表格
      minusRateData: {}, // 电话错误率-资料
      minusRateTable: [], // 电话错误率-表格
      serviceMindsetCountData: {}, // 电话营销意识错误次数-资料
      serviceMindsetCountTable: [], // 电话营销意识错误次数-表格
      serviceTechniqueCountData: {}, // 电话营销技巧错误次数-资料
      serviceTechniqueCountTable: [], // 电话营销技巧错误次数-表格
      serviceAttitudeCountData: {}, // 电话服务态度错误次数-资料
      serviceAttitudeCountTable: [], // 电话服务态度错误次数-表格
      // 图表 tooltip
      tooltip: {
        trigger: "axis",
        axisPointer: {
          type: "cross",
          crossStyle: {
            color: "#909090"
          }
        }
      },
      // 图表 toolbox
      toolbox: {
        feature: {
          dataView: { show: false, readOnly: true },
          magicType: { show: true, type: ["line", "bar"] },
          restore: { show: true },
          saveAsImage: { show: true }
        }
      },
      defaultSeries: [
        {
          name: "",
          type: "bar",
          data: [],
          itemStyle: { normal: { color: "rgba(22, 166, 58, 0.50)" } }
        },
        {
          name: "",
          type: "bar",
          data: [],
          itemStyle: { normal: { color: "rgba(200, 150, 90, 0.50)" } }
        },
        {
          name: "",
          type: "bar",
          data: [],
          itemStyle: { normal: { color: "rgba(12, 80, 231, 0.50)" } }
        },
        {
          name: "",
          type: "bar",
          data: [],
          itemStyle: { normal: { color: "rgba(50, 180, 20, 0.50)" } }
        },
        {
          name: "",
          type: "bar",
          data: [],
          itemStyle: { normal: { color: "rgba(122, 22, 156, 0.50)" } }
        },
        {
          name: "",
          type: "bar",
          data: [],
          itemStyle: { normal: { color: "rgba(123, 99, 10, 0.50)" } }
        },
        {
          name: "",
          type: "bar",
          data: [],
          itemStyle: { normal: { color: "rgba(65, 150, 99, 0.50)" } }
        },
        {
          name: "",
          type: "bar",
          data: [],
          itemStyle: { normal: { color: "rgba(30, 75, 200, 0.50)" } }
        },
        {
          name: "",
          type: "bar",
          data: [],
          itemStyle: { normal: { color: "rgba(50, 23, 77, 0.50)" } }
        },
        {
          name: "",
          type: "bar",
          data: [],
          itemStyle: { normal: { color: "rgba(90, 236, 177, 0.50)" } }
        },
        {
          name: "",
          type: "bar",
          data: [],
          itemStyle: { normal: { color: "rgba(180, 22, 90, 0.50)" } }
        },
        {
          name: "",
          type: "bar",
          data: [],
          itemStyle: { normal: { color: "rgba(210, 80, 97, 0.50)" } }
        },
        {
          name: "",
          type: "bar",
          data: [],
          itemStyle: { normal: { color: "rgba(66, 245, 176, 0.50)" } }
        }
      ],
      dateScale: 1,
      csList: [],
      btnLoading: false
    };
  },
  components: {
    chartTable,
    echart
  },
  methods: {
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
      let totoalDays = moment(endTime).diff(moment(createTime), "days");
      if (totoalDays > self.dateLimit) {
        this.defaultDate();
        self.$alert("开始时间与结束时间间隔不能超过" + self.dateLimit + "天");
      }
      this.transferProductChange();
    },
    checkDialogRange() {
      let self = this;
      let createTime = self.queryMap.times[0];
      let endTime = self.queryMap.times[1];

      let totoalDays = moment(endTime).diff(moment(createTime), "days");
      if (totoalDays > self.dateLimit) {
        this.defaultDate();
        self.$alert("开始时间与结束时间间隔不能超过" + self.dateLimit + "天");
        return false
      }
      this.transferProductChange();
      return true
    },
    /**
     * 预设时间(当前时间往前推n日)
     */
    defaultDate() {
      let startDay= new Date()
      startDay.setDate(startDay.getDate() - this.datedefault)
      //this.$set(this.queryMap.times, 0, new Date(startDay.getFullYear(), startDay.getMonth(), 1, 0, 0, 0))
      if ( this.dateScale==2){
        this.$set(this.queryMap.times,0,new Date(new Date().getFullYear(), new Date().getMonth()-1, 1, 0, 0, 0))
      } else {
        let startDay= new Date()
        startDay.setDate(startDay.getDate() - this.datedefault)
        this.$set(this.queryMap.times,0,new Date(startDay.getFullYear(), startDay.getMonth(), 1, 0, 0, 0))
      }
      this.$set(this.queryMap.times, 1, new Date(startDay.getFullYear(), startDay.getMonth(), startDay.getDate(), 23, 59, 59))
    },
    /**
     * 重置
     */
    reset() {
      this.queryMap = {
        times: [],
        departments: [],
        csType: 1
      };
      this.defaultDate();
      this.initProductIdArray();
    },
    /**
     * 搜寻
     */
    searchData() {
      this.getData();
    },
    lastDay(y, m) {
      return new Date(y, m + 1, 0).getDate();
    },
    /**
     * 更新图表
     */
    async getData() {
      if (!this.checkDialogRange()){
        return
      }
      let self = this;
      self.loading = true;
      self.btnLoading = true;
      let params = {};
      self.queryMap.departments.sort();

      if (self.queryMap.times) {
        if (self.dateScale == 2) {
          let d = new Date(self.queryMap.times[1]);
          d.setDate(self.lastDay(d.getFullYear(), d.getMonth()));
          self.queryMap.times[1] = d;
          d = new Date(self.queryMap.times[0]);
          d.setDate(1);
          self.queryMap.times[0] = d
         }
        const start = moment(self.queryMap.times[0]).format("YYYY-MM-DD 00:00:00");
        const end = moment(self.queryMap.times[1]).format('YYYY-MM-DD 23:59:59')
        params.startTime = start || "";
        params.endTime = end || "";
      }
      if (self.queryMap.departments) {
        params.departments = self.queryMap.departments.toString();
      }

      if (this.csList && this.csList.length > 0) {
        params.agentIds = this.csList.toString();
      }

      if (this.queryMap.csType && this.queryMap.csType != "") {
        params.csType = cloneDeep(this.queryMap.csType);
      }

      params = pickBy(params, (item, key) => {
        return item !== "" && key !== "times";
      });

      let { data: res } = await self.$http.get("system/phoneCsReport/getPhoneCsOverviewReport", { params });

      self.loading = false;
      self.btnLoading = false;

      if (res.success) {
        self.resultDepartments = cloneDeep(self.queryMap.departments);
        // 有效电话量
        if (res.data.effectiveTalkNumberList) {
          let data = self.toFix0(res.data.effectiveTalkNumberList);
          let effectiveTalkNumberDatasets = {
            tooltip: self.tooltip,
            toolbox: self.toolbox,
            legend: {
              data: []
            },
            xAxis: [
              {
                type: "category",
                data: self.resultDepartments,
                axisPointer: {
                  type: "shadow"
                }
              }
            ],
            yAxis: [
              {
                type: "value",
                name: "",
                min: 0,
                // max: 250,
                // interval: 50,
                axisLabel: {
                  formatter: "{value}"
                }
              }
            ],
            series: [
              {
                name: "",
                type: "bar",
                data: [],
                itemStyle: { normal: { color: "rgba(1, 116, 188, 0.50)" } },
                label: { show: false, position: "top" }
              },
              {
                name: "",
                type: "line",
                data: [],
                itemStyle: { normal: { color: "rgba(143, 52, 235, 0.50)" } },
                label: { show: false, position: "top" }
              }
            ]
          };
          effectiveTalkNumberDatasets.series = cloneDeep(
            this.defaultSeries
              .slice(0, res.data.effectiveTalkNumberList.length - 1)
              .concat(effectiveTalkNumberDatasets.series)
          );
          self.effectiveTalkNumberTable = data;
          self.drawChart("effectiveTalkNumberData", data, effectiveTalkNumberDatasets);
        }
        // 人均有效电话量
        if (res.data.averageEffectiveTalkNumberList) {
          let data = self.toFix0(res.data.averageEffectiveTalkNumberList);
          let averageEffectiveTalkNumberDatasets = {
            tooltip: self.tooltip,
            toolbox: self.toolbox,
            legend: {
              data: []
            },
            xAxis: [
              {
                type: "category",
                data: self.resultDepartments,
                axisPointer: {
                  type: "shadow"
                }
              }
            ],
            yAxis: [
              {
                type: "value",
                name: "",
                min: 0,
                // max: 250,
                // interval: 50,
                axisLabel: {
                  formatter: "{value}"
                }
              }
            ],
            series: [
              {
                name: "",
                type: "bar",
                data: [],
                itemStyle: { normal: { color: "rgba(1, 116, 188, 0.50)" } },
                label: { show: false, position: "top" }
              },
              {
                name: "",
                type: "line",
                data: [],
                itemStyle: { normal: { color: "rgba(143, 52, 235, 0.50)" } },
                label: { show: false, position: "top" }
              }
            ]
          };
          averageEffectiveTalkNumberDatasets.series = cloneDeep(
            this.defaultSeries
              .slice(0, res.data.averageEffectiveTalkNumberList.length - 1)
              .concat(averageEffectiveTalkNumberDatasets.series)
          );
          self.averageEffectiveTalkNumberTable = data;
          self.drawChart("averageEffectiveTalkNumberData", data, averageEffectiveTalkNumberDatasets);
        }
        // 人均月通话时长（时）
        if (res.data.averageCostTimePerPersonList) {
          let data = self.toFix0(res.data.averageCostTimePerPersonList);
          let averageCostTimePerPersonDatasets = {
            tooltip: self.tooltip,
            toolbox: self.toolbox,
            legend: {
              data: []
            },
            xAxis: [
              {
                type: "category",
                data: self.resultDepartments,
                axisPointer: {
                  type: "shadow"
                }
              }
            ],
            yAxis: [
              {
                type: "value",
                name: "",
                min: 0,
                // max: 250,
                // interval: 50,
                axisLabel: {
                  formatter: "{value}"
                }
              }
            ],
            series: [
              {
                name: "",
                type: "bar",
                data: [],
                itemStyle: { normal: { color: "rgba(1, 116, 188, 0.50)" } },
                label: { show: false, position: "top" }
              },
              {
                name: "",
                type: "line",
                data: [],
                itemStyle: { normal: { color: "rgba(143, 52, 235, 0.50)" } },
                label: { show: false, position: "top" }
              }
            ]
          };
          averageCostTimePerPersonDatasets.series = cloneDeep(
            this.defaultSeries
              .slice(0, res.data.averageCostTimePerPersonList.length - 1)
              .concat(averageCostTimePerPersonDatasets.series)
          );
          self.averageCostTimePerPersonTable = data;
          self.drawChart("averageCostTimePerPersonData", data, averageCostTimePerPersonDatasets);
        }
        // 平均通话时长（秒）
        if (res.data.averageCostTimeList) {
          let data = self.toFix0(res.data.averageCostTimeList);
          let averageCostTimeDatasets = {
            tooltip: self.tooltip,
            toolbox: self.toolbox,
            legend: {
              data: []
            },
            xAxis: [
              {
                type: "category",
                data: self.resultDepartments,
                axisPointer: {
                  type: "shadow"
                }
              }
            ],
            yAxis: [
              {
                type: "value",
                name: "",
                min: 0,
                // max: 250,
                // interval: 50,
                axisLabel: {
                  formatter: "{value}"
                }
              }
            ],
            series: [
              {
                name: "",
                type: "bar",
                data: [],
                itemStyle: { normal: { color: "rgba(1, 116, 188, 0.50)" } },
                label: { show: false, position: "top" }
              },
              {
                name: "",
                type: "line",
                data: [],
                itemStyle: { normal: { color: "rgba(143, 52, 235, 0.50)" } },
                label: { show: false, position: "top" }
              }
            ]
          };
          averageCostTimeDatasets.series = cloneDeep(
            this.defaultSeries.slice(0, res.data.averageCostTimeList.length - 1).concat(averageCostTimeDatasets.series)
          );
          self.averageCostTimeTable = data;
          self.drawChart("averageCostTimeData", data, averageCostTimeDatasets);
        }
        // 电话呼入接通率
        if (res.data.validCallInputList) {
          let data = self.toFix0(res.data.validCallInputList);
          let validCallInputDatasets = {
            tooltip: self.tooltip,
            toolbox: self.toolbox,
            legend: {
              data: []
            },
            xAxis: [
              {
                type: "category",
                data: self.resultDepartments,
                axisPointer: {
                  type: "shadow"
                }
              }
            ],
            yAxis: [
              {
                type: "value",
                name: "",
                min: 0,
                // max: 250,
                // interval: 50,
                axisLabel: {
                  formatter: "{value}"
                }
              }
            ],
            series: [
              {
                name: "",
                type: "bar",
                data: [],
                itemStyle: { normal: { color: "rgba(1, 116, 188, 0.50)" } },
                label: { show: false, position: "top" }
              },
              {
                name: "",
                type: "line",
                data: [],
                itemStyle: { normal: { color: "rgba(143, 52, 235, 0.50)" } },
                label: { show: false, position: "top" }
              }
            ]
          };
          validCallInputDatasets.series = cloneDeep(
            this.defaultSeries.slice(0, res.data.validCallInputList.length - 1).concat(validCallInputDatasets.series)
          );
          self.validCallInputTable = data;
          self.drawChart("validCallInputData", data, validCallInputDatasets);
        }
        // 电话呼出接通率
        if (res.data.validCallOutputList) {
          let data = self.toFix0(res.data.validCallOutputList);
          let validCallOutputDatasets = {
            tooltip: self.tooltip,
            toolbox: self.toolbox,
            legend: {
              data: []
            },
            xAxis: [
              {
                type: "category",
                data: self.resultDepartments,
                axisPointer: {
                  type: "shadow"
                }
              }
            ],
            yAxis: [
              {
                type: "value",
                name: "",
                min: 0,
                // max: 250,
                // interval: 50,
                axisLabel: {
                  formatter: "{value}"
                }
              }
            ],
            series: [
              {
                name: "",
                type: "bar",
                data: [],
                itemStyle: { normal: { color: "rgba(1, 116, 188, 0.50)" } },
                label: { show: false, position: "top" }
              },
              {
                name: "",
                type: "line",
                data: [],
                itemStyle: { normal: { color: "rgba(143, 52, 235, 0.50)" } },
                label: { show: false, position: "top" }
              }
            ]
          };
          validCallOutputDatasets.series = cloneDeep(
            this.defaultSeries.slice(0, res.data.validCallOutputList.length - 1).concat(validCallOutputDatasets.series)
          );
          self.validCallOutputTable = data;
          self.drawChart("validCallOutputData", data, validCallOutputDatasets);
        }
        // 电话总接通率
        if (res.data.validCallTotalList) {
          let data = self.toFix0(res.data.validCallTotalList);
          let validCallTotalDatasets = {
            tooltip: self.tooltip,
            toolbox: self.toolbox,
            legend: {
              data: []
            },
            xAxis: [
              {
                type: "category",
                data: self.resultDepartments,
                axisPointer: {
                  type: "shadow"
                }
              }
            ],
            yAxis: [
              {
                type: "value",
                name: "",
                min: 0,
                // max: 250,
                // interval: 50,
                axisLabel: {
                  formatter: "{value}"
                }
              }
            ],
            series: [
              {
                name: "",
                type: "bar",
                data: [],
                itemStyle: { normal: { color: "rgba(1, 116, 188, 0.50)" } },
                label: { show: false, position: "top" }
              },
              {
                name: "",
                type: "line",
                data: [],
                itemStyle: { normal: { color: "rgba(143, 52, 235, 0.50)" } },
                label: { show: false, position: "top" }
              }
            ]
          };
          validCallTotalDatasets.series = cloneDeep(
            this.defaultSeries.slice(0, res.data.validCallTotalList.length - 1).concat(validCallTotalDatasets.series)
          );
          self.validCallTotalTable = data;
          self.drawChart("validCallTotalData", data, validCallTotalDatasets);
        }
        // 电话嘉奖率
        if (res.data.plusRateList) {
          let data = self.toFix0(res.data.plusRateList);
          let plusRateDatasets = {
            tooltip: self.tooltip,
            toolbox: self.toolbox,
            legend: {
              data: []
            },
            xAxis: [
              {
                type: "category",
                data: self.resultDepartments,
                axisPointer: {
                  type: "shadow"
                }
              }
            ],
            yAxis: [
              {
                type: "value",
                name: "",
                min: 0,
                // max: 250,
                // interval: 50,
                axisLabel: {
                  formatter: "{value}"
                }
              }
            ],
            series: [
              {
                name: "",
                type: "bar",
                data: [],
                itemStyle: { normal: { color: "rgba(1, 116, 188, 0.50)" } },
                label: { show: false, position: "top" }
              },
              {
                name: "",
                type: "line",
                data: [],
                itemStyle: { normal: { color: "rgba(143, 52, 235, 0.50)" } },
                label: { show: false, position: "top" }
              }
            ]
          };
          plusRateDatasets.series = cloneDeep(
            this.defaultSeries.slice(0, res.data.plusRateList.length - 1).concat(plusRateDatasets.series)
          );
          self.plusRateTable = data;
          self.drawChart("plusRateData", data, plusRateDatasets);
        }
        // 电话错误率
        if (res.data.minusRateList) {
          let data = self.toFix0(res.data.minusRateList);
          let minusRateDatasets = {
            tooltip: self.tooltip,
            toolbox: self.toolbox,
            legend: {
              data: []
            },
            xAxis: [
              {
                type: "category",
                data: self.resultDepartments,
                axisPointer: {
                  type: "shadow"
                }
              }
            ],
            yAxis: [
              {
                type: "value",
                name: "",
                min: 0,
                // max: 250,
                // interval: 50,
                axisLabel: {
                  formatter: "{value}"
                }
              }
            ],
            series: [
              {
                name: "",
                type: "bar",
                data: [],
                itemStyle: { normal: { color: "rgba(1, 116, 188, 0.50)" } },
                label: { show: false, position: "top" }
              },
              {
                name: "",
                type: "line",
                data: [],
                itemStyle: { normal: { color: "rgba(143, 52, 235, 0.50)" } },
                label: { show: false, position: "top" }
              }
            ]
          };
          minusRateDatasets.series = cloneDeep(
            this.defaultSeries.slice(0, res.data.minusRateList.length - 1).concat(minusRateDatasets.series)
          );
          self.minusRateTable = data;
          self.drawChart("minusRateData", data, minusRateDatasets);
        }
        // 电话营销意识错误次数
        if (res.data.serviceMindsetCountList) {
          let data = self.toFix0(res.data.serviceMindsetCountList);
          let serviceMindsetCountDatasets = {
            tooltip: self.tooltip,
            toolbox: self.toolbox,
            legend: {
              data: []
            },
            xAxis: [
              {
                type: "category",
                data: self.resultDepartments,
                axisPointer: {
                  type: "shadow"
                }
              }
            ],
            yAxis: [
              {
                type: "value",
                name: "",
                min: 0,
                // max: 250,
                // interval: 50,
                axisLabel: {
                  formatter: "{value}"
                }
              }
            ],
            series: [
              {
                name: "",
                type: "bar",
                data: [],
                itemStyle: { normal: { color: "rgba(1, 116, 188, 0.50)" } },
                label: { show: false, position: "top" }
              },
              {
                name: "",
                type: "line",
                data: [],
                itemStyle: { normal: { color: "rgba(143, 52, 235, 0.50)" } },
                label: { show: false, position: "top" }
              }
            ]
          };
          serviceMindsetCountDatasets.series = cloneDeep(
            this.defaultSeries
              .slice(0, res.data.serviceMindsetCountList.length - 1)
              .concat(serviceMindsetCountDatasets.series)
          );
          self.serviceMindsetCountTable = data;
          self.drawChart("serviceMindsetCountData", data, serviceMindsetCountDatasets);
        }
        // 电话营销技巧错误次数
        if (res.data.serviceTechniqueCountList) {
          let data = self.toFix0(res.data.serviceTechniqueCountList);
          let serviceTechniqueCountDatasets = {
            tooltip: self.tooltip,
            toolbox: self.toolbox,
            legend: {
              data: []
            },
            xAxis: [
              {
                type: "category",
                data: self.resultDepartments,
                axisPointer: {
                  type: "shadow"
                }
              }
            ],
            yAxis: [
              {
                type: "value",
                name: "",
                min: 0,
                // max: 250,
                // interval: 50,
                axisLabel: {
                  formatter: "{value}"
                }
              }
            ],
            series: [
              {
                name: "",
                type: "bar",
                data: [],
                itemStyle: { normal: { color: "rgba(1, 116, 188, 0.50)" } },
                label: { show: false, position: "top" }
              },
              {
                name: "",
                type: "line",
                data: [],
                itemStyle: { normal: { color: "rgba(143, 52, 235, 0.50)" } },
                label: { show: false, position: "top" }
              }
            ]
          };
          serviceTechniqueCountDatasets.series = cloneDeep(
            this.defaultSeries
              .slice(0, res.data.serviceTechniqueCountList.length - 1)
              .concat(serviceTechniqueCountDatasets.series)
          );
          self.serviceTechniqueCountTable = data;
          self.drawChart("serviceTechniqueCountData", data, serviceTechniqueCountDatasets);
        }
        // 电话服务态度错误次数
        if (res.data.serviceAttitudeCountList) {
          let data = self.toFix0(res.data.serviceAttitudeCountList);
          let serviceAttitudeCountDatasets = {
            tooltip: self.tooltip,
            toolbox: self.toolbox,
            legend: {
              data: []
            },
            xAxis: [
              {
                type: "category",
                data: self.resultDepartments,
                axisPointer: {
                  type: "shadow"
                }
              }
            ],
            yAxis: [
              {
                type: "value",
                name: "",
                min: 0,
                // max: 250,
                // interval: 50,
                axisLabel: {
                  formatter: "{value}"
                }
              }
            ],
            series: [
              {
                name: "",
                type: "bar",
                data: [],
                itemStyle: { normal: { color: "rgba(1, 116, 188, 0.50)" } },
                label: { show: false, position: "top" }
              },
              {
                name: "",
                type: "line",
                data: [],
                itemStyle: { normal: { color: "rgba(143, 52, 235, 0.50)" } },
                label: { show: false, position: "top" }
              }
            ]
          };
          serviceAttitudeCountDatasets.series = cloneDeep(
            this.defaultSeries
              .slice(0, res.data.serviceAttitudeCountList.length - 1)
              .concat(serviceAttitudeCountDatasets.series)
          );
          self.serviceAttitudeCountTable = data;
          self.drawChart("serviceAttitudeCountData", data, serviceAttitudeCountDatasets);
        }
      } else {
        self.$message.error({
          message: res.data.errorMsg,
          type: "error"
        });
      }
    },
    // 更新图表 - 长条图
    drawChart(chartData, newData, sets) {
      let self = this;
      let labels = cloneDeep(self.queryMap.departments);
      labels.push("平均值");
      let list = cloneDeep(self.queryMap.departments);
      list.push("mean");
      self[chartData] = null;
      let newChartData = sets;
      if (newData) {
        newChartData.xAxis[0].data = labels;
        const averageIndex = newData.length - 1;
        newData.forEach((a, index) => {
          let data = [];
          list.forEach((b) => {
            data.push(a[b]);
          });
          newChartData.series[index].data = data;
          newChartData.series[index].name = a.title;
          newChartData.legend.data[index] = a.title;
          // 调整平均值的图表类型及显示标签
          if (index === averageIndex) {
            newChartData.series[index].type = "line";
            newChartData.series[index].label.show = true;
          }
        });
        self[chartData] = newChartData;
      } else {
        self[chartData] = {};
        self.$message.error({
          message: chartData + "无图表资料",
          type: "error"
        });
      }
    },
    // 数值取小数点后两位，四舍五入
    toFix2(array) {
      if (array[0]) {
        let newArray = cloneDeep(array);
        newArray.forEach((a) => {
          this.queryMap.departments.forEach((b) => {
            a[b] = parseFloat(Number(a[b]).toFixed(2));
          });
          a["mean"] = parseFloat(Number(a["mean"]).toFixed(2));
        });
        return newArray;
      }
      return [];
    },
    // 数值取整数，四舍五入
    toFix0(array) {
      if (array[0]) {
        let newArray = cloneDeep(array);
        newArray.forEach((a) => {
          this.queryMap.departments.forEach((b) => {
            a[b] = parseFloat(Number(a[b]).toFixed(0));
          });
          a["mean"] = parseFloat(Number(a["mean"]).toFixed(0));
        });
        return newArray;
      }
      return [];
    },
    async getAllDepartments() {
      const { data: res } = await this.$http.get("/system/overviewReport/allDepartments");
      return new Promise((resolve, rej) => {
        resolve(res.data);
      });
    },
    setDateScale(type) {
      this.dateScale = type;
      //http://jira.pai9.net:8080/browse/DS-1249
      if (type==2){
        this.$set(this.queryMap.times,0,new Date(new Date().getFullYear(), new Date().getMonth()-1, 1, 0, 0, 0))
      } else {
        let startDay= new Date()
        startDay.setDate(startDay.getDate() - this.datedefault)
        this.$set(this.queryMap.times,0,new Date(startDay.getFullYear(), startDay.getMonth(), 1, 0, 0, 0))
      }
    },
    async transferProductChange() {
      this.btnLoading = true;
      let params = {
        department: this.queryMap.department,
        startTime: moment(this.queryMap.times[0]).format("YYYY-MM-DD HH:mm:ss"),
        endTime: moment(this.queryMap.times[1]).format("YYYY-MM-DD HH:mm:ss"),
        csType: this.queryMap.csType
      };
      const { data: res } = await this.$http.get("system/agentServiceSound/getCs", { params });
      if (!res.success) {
        this.$message.error("获取客服数据失败:" + res.data.errorMsg);
      } else {
        this.csList = res.data.map((i) => i.rawId);
      }
      this.btnLoading = false;
    },
    async initData() {
      this.defaultDate();
      this.loading = true;
      await this.transferProductChange();
      this.getData();
      this.loading = false;
    },
    initProductIdArray(){
      let userInfo = this.$store.state.userInfo;
      let productIdArray = userInfo.productIdArray;
      if(userInfo.type === 4){
        let perms = userInfo.perms;
        if(null!= perms && perms.length >0){
          let arr = [];
          for(let i in perms){
            let productId = perms[i].replace('phoneCsOverview:','');
            for(let i in constants.DEPARTMENT_LIST){
              if(constants.DEPARTMENT_LIST[i].key === productId){
                arr.push(productId);
                break;
              }
            }
          }
         this.queryMap.departments = productIdArray.concat(arr.filter(s=>productIdArray.indexOf(s)===-1));
          this.departmentList = productIdArray.concat(arr.filter(s=>productIdArray.indexOf(s)===-1));
        }else {
          this.departmentList = [];
        }
      } else {
        this.queryMap.departments = productIdArray;
        this.departmentList = productIdArray;
      }



    }
  },
  created() {
    this.initProductIdArray();
    this.initData();
  },
  mounted() {}
};
</script>

<style scoped>
.chart-table {
  margin-top: 20px;
  margin-right: 2px;
  margin-bottom: 40px;
  width: 100%;
}
.sub-title {
  font-size: 16px;
  margin-top: 30px;
  margin-bottom: 10px;
  font-weight: bold;
  text-align: center;
}
.block-fliter {
  margin-bottom: 20px;
}
.el-checkbox {
  margin-right: 10px;
}
.chartBar {
  padding-left: 60px;
  width: calc(100% - 60px);
}
.el-select,
.el-input {
  width: 128px;
}
.datePicker {
  width: 335px;
}
</style>
