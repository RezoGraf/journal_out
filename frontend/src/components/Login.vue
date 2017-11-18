<template>
  <div id="general">
    <b-container class="col-md-3">
    <br />
      <b-card>
      <div class="text-center"> <h2>Вакцинация</h2></div>
      <b-row class="my-1">
        <b-col>
          Введите имя пользователя и пароль
          <b-form-input v-model="name" type="text" placeholder="Имя пользователя"></b-form-input>
        </b-col>
      </b-row>
      <b-row class="my-1">
        <b-col>
          <b-form-input v-model="pass" type="password" v-on:keyup.enter="console.log('12312312')" placeholder="Пароль"></b-form-input>
        </b-col>
      </b-row>
      <b-row class="my-1">
        <b-col>
          <b-button type="text" variant="primary" v-on:click="login">Вход</b-button>
        </b-col>
      </b-row>
      </b-card>
      <!--<b-alert v-if="this.errauth" variant="success" show>Success Alert</b-alert>-->
    </b-container>
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
    data: function () {
      return {
        name: null,
        pass: null,
        errauth: false
      }
    },
    mounted: function () {
      Auth.checkToken()
    },
    methods: {
      login: function () {
        Auth.login(this.name, this.pass)
        if (Auth.errauth) {
          this.errauth = true
        }
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
