(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-5e85ad54"],{2423:function(t,e,r){"use strict";r.d(e,"a",(function(){return i})),r.d(e,"b",(function(){return c})),r.d(e,"c",(function(){return l})),r.d(e,"d",(function(){return n})),r.d(e,"e",(function(){return s}));var a=r("b775");function i(t){return Object(a["a"])({url:"/lego/article/createarticle",method:"post",data:t})}function c(t){return Object(a["a"])({url:"/lego/article/deletearticle",method:"post",data:t})}function l(t){return Object(a["a"])({url:"/lego/article/getarticle",method:"post",data:t})}function n(t){return Object(a["a"])({url:"/lego/article/getauthoIrdarticles",method:"post",data:t})}function s(t){return Object(a["a"])({url:"/lego/article/getlastarticles",method:"post",data:t})}},"558f":function(t,e,r){"use strict";r.r(e);var a=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("v-container",{staticClass:"d-flex justify-center mt-6"},[r("v-card",{attrs:{width:"100%",height:"100%",flat:""}},[r("v-card-text",[r("v-row",{attrs:{justify:"center"}},[r("v-col",{attrs:{cols:"3"}},[r("v-card",{attrs:{height:"830px"}},[r("v-card-title",[t._v("文章列表")]),r("v-divider"),r("v-card-text",[t.articles.length>0?r("v-responsive",{staticClass:"overflow-y-auto",attrs:{"max-height":"720"}},[t._l(t.articles,(function(e,a){return[r("v-card",{key:a,staticClass:" pa-2",attrs:{outlined:""},on:{click:function(e){return t.getarticlehandle(a)}}},[r("v-card",{staticClass:"text-h5 mb-2 transparent",attrs:{flat:""}},[t._v(" "+t._s(e.Title)+" ")]),r("v-card",{staticClass:"mt-1 text-justify text--disabled transparent",attrs:{width:"100%",flat:""}},[t._v(" "+t._s(t._f("ellipsis")(e.Content))+" ")]),r("v-divider",{staticClass:"mt-2"}),r("v-card",{staticClass:"d-flex justify-end mt-1 transparent",attrs:{flat:""}},[r("v-card",{staticClass:"text--disabled transparent",attrs:{flat:""}},[t._v(t._s(t._f("formatDate")(e.CreationTime)))])],1)],1)]}))],2):r("v-card",{staticClass:"text-center",attrs:{height:"720px",flat:""}},[r("v-card-text",{staticClass:"text--disabled"},[t._v(" 您当前还没有文章")])],1)],1)],1),r("v-btn",{staticClass:"mt-2 red",attrs:{dark:"",block:""},on:{click:t.restarteditor}},[t._v("添加新的文章")]),r("v-btn",{staticClass:"mt-2 green",attrs:{disabled:0==t.currarticle.Id,block:""},on:{click:t.deleteArticle}},[t._v("删除文章")])],1),r("v-col",{attrs:{cols:"9"}},[r("v-card",[r("v-card-title",[t._v("编辑文章")]),r("v-divider"),r("v-card-text",[r("v-text-field",{attrs:{label:"Title"},model:{value:t.currarticle.Title,callback:function(e){t.$set(t.currarticle,"Title",e)},expression:"currarticle.Title"}}),r("mavon-editor",{ref:"md",on:{imgAdd:t.$imgAdd,imgDel:t.$imgDel},model:{value:t.currarticle.Content,callback:function(e){t.$set(t.currarticle,"Content",e)},expression:"currarticle.Content"}})],1),r("v-card-actions",{staticClass:"d-flex justify-end"},[r("v-btn",{attrs:{rounded:""},on:{click:t.saveearticle}},[t._v("保存")]),r("v-btn",{staticClass:"text-h6  ml-6 white--text",attrs:{color:"error",rounded:""},on:{click:t.releasearticle}},[t._v("发布")])],1)],1)],1)],1)],1)],1)],1)},i=[],c=(r("a9e3"),r("fb6a"),r("a434"),r("ac1f"),r("5319"),r("2423")),l=r("c1df"),n=r.n(l),s={props:{articleid:{type:Number,default:0}},data:function(){return{articles:[],currarticle:{Id:0,Title:"",Content:"",ShortContent:"",Images:[],CreationTime:0},img_file:{}}},filters:{ellipsis:function(t){return t?t.length>60?t.slice(0,60)+"...":t:""},formatDate:function(t){var e=new Date(1e3*t);return n()(e).format("YYYY-MM-DD HH:mm:ss")}},created:function(){var t=this;Object(c["d"])(null).then((function(e){t.articles=e.data}))},methods:{$imgAdd:function(t,e){console.log("添加图片"+e),this.img_file[t]=e},$imgDel:function(t){delete this.img_file[t]},uploadimg:function(t){var e=new FormData;for(var r in this.img_file)e.append(r,this.img_file[r]);axios({url:"server url",method:"post",data:e,headers:{"Content-Type":"multipart/form-data"}}).then((function(t){for(var e in t)$vm.$img2Url(e[0],e[1])}))},getarticlehandle:function(t){this.currarticle=this.articles[t]},saveearticle:function(){var t=this;this.getmdimage(),this.getshortContent(),Object(c["a"])({ArticleId:this.currarticle.Id,Title:this.currarticle.Title,Content:this.currarticle.Content,ShortContent:this.currarticle.ShortContent,Images:this.currarticle.Images}).then((function(e){e.data.Id!=t.currarticle.Id&&t.articles.push(e.data);var r=e.message;t.$message.success(r)}))},releasearticle:function(){},deleteArticle:function(){var t=this;Object(c["b"])({ArticleId:this.currarticle.Id}).then((function(e){for(var r=0,a=t.articles.length;r<a;r++)t.articles[r].Id==e.data.Id&&t.articles.splice(r,1);t.currarticle={Id:0,Title:"",Content:"",CreationTime:0};var i=e.message;t.$message.success(i)}))},restarteditor:function(){this.currarticle={Id:0,Title:"",Content:"",CreationTime:0}},getmdimage:function(){var t,e=/!\[(.*?)\]\((.*?)\)/gm;this.currarticle.Images=[];while(null!==(t=e.exec(this.currarticle.Content)))this.currarticle.Images.push(t[2]);console.log(this.currarticle.Images)},getshortContent:function(){if(""!=this.currarticle.Content){var t=this.currarticle.Content.replace(/(\*\*|__)(.*?)(\*\*|__)/g,"").replace(/\!\[[\s\S]*?\]\([\s\S]*?\)/g,"").replace(/\[[\s\S]*?\]\([\s\S]*?\)/g,"").replace(/<\/?.+?\/?>/g,"").replace(/(\*)(.*?)(\*)/g,"").replace(/`{1,2}[^`](.*?)`{1,2}/g,"").replace(/```([\s\S]*?)```[\s]*/g,"").replace(/\~\~(.*?)\~\~/g,"").replace(/[\s]*[-\*\+]+(.*)/g,"").replace(/[\s]*[0-9]+\.(.*)/g,"").replace(/(#+)(.*)/g,"").replace(/(>+)(.*)/g,"").replace(/\r\n/g,"").replace(/\n/g,"").replace(/\s/g,"");this.currarticle.ShortContent=t.slice(0,155)}}}},o=s,d=r("2877"),u=r("6544"),f=r.n(u),h=r("8336"),v=r("b0afa"),g=r("99d9"),m=r("62ad"),p=r("a523"),C=r("ce7e"),b=r("6b53"),_=r("0fd9"),x=r("8654"),I=Object(d["a"])(o,a,i,!1,null,null,null);e["default"]=I.exports;f()(I,{VBtn:h["a"],VCard:v["a"],VCardActions:g["a"],VCardText:g["b"],VCardTitle:g["c"],VCol:m["a"],VContainer:p["a"],VDivider:C["a"],VResponsive:b["a"],VRow:_["a"],VTextField:x["a"]})}}]);
//# sourceMappingURL=chunk-5e85ad54.7accc46a.js.map