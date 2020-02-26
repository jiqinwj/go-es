<template>
    <div>
      <div>
          <dl class="sdt">
              <dt>
                  出版社
              </dt>
              <dd v-for="item in pressList">
                  <a href="#" :class="searchModel.book_press === item?'select':''" @click="setPress(item)">{{item}}</a>
              </dd>
          </dl>

          <dl class="sdt">
              <dt>
                  价格范围:
              </dt>
              <dd>
                  <el-input    class="numtext" v-model.number="searchModel.book_price1_start" placeholder="输入数字"></el-input>
                  <span>~~</span>
                  <el-input    class="numtext" v-model.number="searchModel.book_price1_end" placeholder="输入数字"></el-input>
              </dd>
          </dl>
          <dl class="sdt">
              <dt>
                  全文检索
              </dt>
              <dd class="row">
                  <el-input  v-model="searchModel.book_name"   class="search" placeholder="请输入内容"></el-input>
                  <el-button type="primary"  @click="search" >搜索</el-button>
              </dd>
          </dl>
          <dl class="sdt">
              <dt>
                  排序方式:
              </dt>
              <dd class="row">
                  <a   :class="searchModel.OrderSet.score?'select':''" @click="searchModel.OrderSet.score=!searchModel.OrderSet.score"  >综合</a>
                  <el-dropdown @command="(command)=>{searchModel.OrderSet.price_order=command}" >
                      <span class="el-dropdown-link">
                        {{PriceSelect[searchModel.OrderSet.price_order]}}<i class="el-icon-arrow-down el-icon--right"></i>
                      </span>
                      <el-dropdown-menu slot="dropdown">
                          <template v-for="(item,index) in PriceSelect">
                              <el-dropdown-item :command="index">{{item}}</el-dropdown-item>
                          </template>
                      </el-dropdown-menu>
                  </el-dropdown>

              </dd>
          </dl>
      </div>
        <div style="width: 95%;margin:30px auto">
            <el-table
                    :data="bookList"
                    border
                    style="width: 100%">
                <el-table-column
                        prop="BookID"
                        label="ID"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="BookName"
                        label="商品名称"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="BookIntr"
                        label="商品简介"
                        width="280">
                </el-table-column>
                <el-table-column
                        prop="BookPrice1"
                        label="价格">
                </el-table-column>
                <el-table-column
                        prop="BookPress"
                        label="出版社">
                </el-table-column>
                <el-table-column
                        prop="BookAuthor"
                        label="作者">
                </el-table-column>
            </el-table>
        </div>

    </div>

</template>
<script>
   module.exports ={
       data(){
           return {
               bookList: [],
               pressList: [],
               searchModel: {
                   book_name: '',
                   book_press: '',
                   book_price1_start:0,
                   book_price1_end:0,
                   OrderSet:{
                       score:true, // 默认按socre
                       price_order:0,// 0就是不限价格 1代表从低到高，2代表从高到低
                   }
               },
               PriceSelect:["价格不限","价格从低到高","价格从高到低"]
           }
       },
       created(){
            this.loadPress()  // 加载出版社列表
       },
       methods: {
           async loadPress(){
               const response = await axios.get("/helper/press")
               const { result } = response.data
               this.pressList = result
           },
           async search(){
               console.log(this.searchModel)
               const response = await axios.post("/books/search",this.searchModel)
               const { result } = response.data
               this.bookList = result
           },
           setPress(press){
               if (this.searchModel.book_press === press ){
                   this.searchModel.book_press = ''
               }else {
                   this.searchModel.book_press = press
               }

           }
       }
   }
</script>
<style>
  .sdt{margin: 10px auto;width:90%;border-radius: 5px;display: block;float:left;margin-left: 50px}
  .sdt dt{width:100%;display: block;color:#3A7B43;font-size:16px;font-weight: bold;margin-bottom: 10px}
  .sdt dd{float:left;margin:0 auto;text-indent:1em}
  .sdt .row{width:100%;}
  .sdt dd .search{width:50%;}
  .sdt dd a{color: #3a8ee6}
  .sdt dd a:hover{background: #eb5975;color:#fff}
  .sdt dd .select{background: #eb5975;color:#fff}
  .sdt dd .numtext{width:100px}
  .sdt dd span{margin: 0 auto}
    a{cursor:pointer}
</style>