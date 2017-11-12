<template>
  <div id="general">
    <input type="submit" value="login" v-on:click="login">
  </div>
</template>

<script>
  import Vue from 'vue'
  import BootstrapVue from 'bootstrap-vue'
  import 'bootstrap/dist/css/bootstrap.css'
  import 'bootstrap-vue/dist/bootstrap-vue.css'
  import axios from 'axios'
//  import router from 'vue-router'
  import {Auth} from '../router/index.js'

  Vue.use(BootstrapVue)
  export default {
    name: 'general',
    data () {
      return {
        items: [],
        tabIndex: 0,
        input_find: '',
        Fio: '',
        Pasport: '',
        DateRogd: '',
        Pol: '',
        msg: 'Welcome to Your Vue.js App',
        selected: null,
        options: [
          {value: null, text: 'Кем выдан'},
          {value: 'a', text: 'УВД Ленинского района'},
          {value: 'b', text: 'УВД Кировского района'},
          {value: 'c', text: 'УВД Центрального района'}
        ],
        fieldsDift: [
          {key: 'DATEPRIV', label: 'Дата прививки', class: 'text-justify col-xs-8'},
          {key: 'PREPARAT', label: 'Препарат', class: 'text-justify'},
          {key: 'PRIVIVKI', label: 'Прививки', class: 'text-center'},
          {key: 'SERIA', label: 'Серия', class: 'text-justify'},
          {key: 'DOZA', label: 'Доза', class: 'text-center'},
          {key: 'MEDOTVOD', label: 'Медотвод', class: 'text-center'}]
      }
    },
    mounted: function () {
      Auth.checkToken()
    },
    methods: {
      login: function () {
        Auth.login()
      },
      startAutoUpdate: function () {
        this.timer = setInterval(this.GetModeVrach, 600000)
      },
      cancelAutoUpdate: function () {
        clearInterval(this.timer)
      },
      GetPatient () {
        const data = new FormData()
        var ss = this
        data.append('patient', ss.input_find)
        axios.post('http://192.168.1.76:8084/FindPatientInArena', data, {responeType: 'application/json'})
          .then(function (response) {
            var data = JSON.parse(JSON.stringify(response.data))
//          ss.items = data
//          console.log(ss.items)
            ss.Fio = data[0].Fio
            ss.Pasport = data[0].Pasport
            ss.DateRogd = data[0].DateRogd
            ss.Pol = data[0].Pol
            console.log(ss.Fio)
//          ss.saveJSON(ss.items)
          })
          .catch(function (error) {
            console.log(error)
//          ss.items = JSON.parse(localStorage.getItem('myObj'))
          })
      }
    }
  }
</script>
