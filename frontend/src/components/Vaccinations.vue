<template>
  <div class="general">
      <b-container>
      <br>
        <b-row class="mb-4 text-right">
          <b-col cols="3">
            <b-form-input v-model="input_find" placeholder="№ Карты"></b-form-input>
          </b-col>
          <b-col cols="1">
            <b-button variant="primary" @click="this.GetPatient">Найти</b-button>
          </b-col>
          <b-col md="8">
            <b-button variant="primary" @click="this.logout">Выйти</b-button>
          </b-col>
        </b-row>
        <b-row class="mb1 text-center">
        </b-row>

        <b-row>
          <b-col>

            <!-- Tabs with card integration -->
            <!--<b-col class="mb-1 text-center">Вакцинация</b-col>-->
            <b-card no-body>
              <!--<b-col-md-4>-->
                <b-card>
                <b> Ф.И.О: </b> {{Fio}}<br>
                <b>Паспорт:</b> {{Pasport}}<br>
                <b>Дата рождения:</b> {{DateRogd}} <br>
                <b>Пол:</b> {{Pol}}
                </b-card>
              <!--</b-col-md-4>-->
              <b-tabs small card ref="tabs" v-model="tabIndex">
 <!--Дифтерия  Начало таблицы          -->
                <b-tab title="Дифтерия">
                  <b-table striped hover
                           :items="items"
                           :fields="fieldsDift">
                    <template slot="DATEPRIV" scope="data">
                      {{data.item.DatePriv}}
                    </template>
                    <template slot="PREPARAT" scope="data">
                      {{data.item.Preparat}}
                    </template>
                    <template slot="PRIVIVKI" scope="data">
                      {{data.item.Privivki}}
                    </template>
                    <template slot="DOZA" scope="data">
                      {{data.item.Doza}}
                    </template>
                    <template slot="MEDOTVOD" scope="data">
                      {{data.item.MedOtvod}}
                    </template>
                  </b-table>
   <!--Дифтерия  конец таблицы          -->
                  <b-button variant="primary" size="sm" v-b-modal.ModalDifteria>Добавить</b-button>
                  <b-modal id="ModalDifteria">
                    <b-container>
                      <b-row>Дата вакцинации: <b-col align-self="center"><b-input type="date"></b-input></b-col></b-row> <br />
                      <b-row><b-input type="text" placeholder="Препарат"></b-input></b-row><br />
                      <b-row><b-input type="text" placeholder="Прививки"></b-input></b-row><br />
                      <b-row><b-input type="text" placeholder="Серия"></b-input></b-row><br />
                      <b-row><b-input type="text" placeholder="Доза"></b-input></b-row><br />
                      <b-row><b-input type="text" placeholder="Медотвод"></b-input></b-row>
                    </b-container>
                  </b-modal>
                </b-tab>
                <b-tab title="Кл. энцефалит">
                  I'm the second tab
                  <b-card>I'm the card in tab</b-card>
                </b-tab>
                <b-tab title="Корь">
                  I'm the second tab
                  <b-card>I'm the card in tab</b-card>
                </b-tab>
                <b-tab title="Краснуха">
                  I'm the second tab
                  <b-card>I'm the card in tab</b-card>
                </b-tab>
                <b-tab title="Гепатит В">
                  I'm the second tab
                  <b-card>I'm the card in tab</b-card>
                </b-tab>
                <b-tab title="Грипп">
                  I'm the second tab
                  <b-card>I'm the card in tab</b-card>
                </b-tab>
              </b-tabs>
            </b-card>

          </b-col>

        </b-row>
      </b-container>
  </div>
</template>

<script>
import Vue from 'vue'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import axios from 'axios'
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
  methods: {
    logout: function () {
      Auth.logout()
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
      axios.post('http://localhost:8084/jwt/FindPatientInArena', data, {responeType: 'application/json'})
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
