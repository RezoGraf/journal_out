<template>
  <div class="general">
      <b-container>
      <br>
        <b-row class="mb-4 text-right">
          <b-col cols="3">
            <b-form-input v-model="input_find" placeholder="№ Карты"></b-form-input>
          </b-col>
          <b-col cols="1">
            <b-button variant="primary" @click="this.GetPatient" :disabled="this.input_find == ''">Найти</b-button>
          </b-col>
          <b-col md="8">
            {{nameUser.name}} &nbsp;
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
                <b-tab title="Дифтерия" @click="setItemsDefaultVACCINATIONS('АДСМ')">
                  <b-table striped hover
                           :items="itemsVACCINATIONS"
                           :fields="fieldsDift">
                    <template slot="DATEPRIV" scope="data">
                      {{data.item.date}}
                    </template>
                    <template slot="PREPARAT" scope="data">
                      {{data.item.preparat}}
                    </template>
                    <template slot="SERIA" scope="data">
                      {{data.item.seria}}
                    </template>
                    <template slot="DOZA" scope="data">
                      {{data.item.doza}}
                    </template>
                  </b-table>
   <!--Дифтерия  конец таблицы          -->



                </b-tab>
                  <b-tab title="Клещевой энцефалит" @click="setItemsDefaultVACCINATIONS('Клещевой энцефалит')">
                    <b-table striped hover
                             :items="itemsVACCINATIONS"
                             :fields="fieldsKLESH">
                      <template slot="DATEPRIV" scope="data">
                        {{data.item.date}}
                      </template>
                      <template slot="PREPARAT" scope="data">
                        {{data.item.preparat}}
                      </template>
                      <template slot="SERIA" scope="data">
                        {{data.item.seria}}
                      </template>
                      <template slot="DOZA" scope="data">
                        {{data.item.doza}}
                      </template>
                    </b-table>
                </b-tab>
                <b-tab title="Корь" @click="setItemsDefaultVACCINATIONS('Корь')">
                  <b-table striped hover
                           :items="itemsVACCINATIONS"
                           :fields="fieldsKLESH">
                    <template slot="DATEPRIV" scope="data">
                      {{data.item.date}}
                    </template>
                    <template slot="PREPARAT" scope="data">
                      {{data.item.preparat}}
                    </template>
                    <template slot="SERIA" scope="data">
                      {{data.item.seria}}
                    </template>
                    <template slot="DOZA" scope="data">
                      {{data.item.doza}}
                    </template>
                  </b-table>
                </b-tab>
                <b-tab title="Краснуха" @click="setItemsDefaultVACCINATIONS('Краснуха')">
                  <b-table striped hover
                           :items="itemsVACCINATIONS"
                           :fields="fieldsKLESH">
                    <template slot="DATEPRIV" scope="data">
                      {{data.item.date}}
                    </template>
                    <template slot="PREPARAT" scope="data">
                      {{data.item.preparat}}
                    </template>
                    <template slot="SERIA" scope="data">
                      {{data.item.seria}}
                    </template>
                    <template slot="DOZA" scope="data">
                      {{data.item.doza}}
                    </template>
                  </b-table>
                </b-tab>
                <b-tab title="Гепатит В" @click="setItemsDefaultVACCINATIONS('Гепатит В')">
                  <b-table striped hover
                           :items="itemsVACCINATIONS"
                           :fields="fieldsKLESH">
                    <template slot="DATEPRIV" scope="data">
                      {{data.item.date}}
                    </template>
                    <template slot="PREPARAT" scope="data">
                      {{data.item.preparat}}
                    </template>
                    <template slot="SERIA" scope="data">
                      {{data.item.seria}}
                    </template>
                    <template slot="DOZA" scope="data">
                      {{data.item.doza}}
                    </template>
                  </b-table>
                </b-tab>
                <b-tab title="Грипп" @click="setItemsDefaultVACCINATIONS('Грипп')">
                  <b-table striped hover
                           :items="itemsVACCINATIONS"
                           :fields="fieldsKLESH">
                    <template slot="DATEPRIV" scope="data">
                      {{data.item.date}}
                    </template>
                    <template slot="PREPARAT" scope="data">
                      {{data.item.preparat}}
                    </template>
                    <template slot="SERIA" scope="data">
                      {{data.item.seria}}
                    </template>
                    <template slot="DOZA" scope="data">
                      {{data.item.doza}}
                    </template>
                  </b-table>
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
import * as controllers from '../../controllers/Controllers.js'

Vue.use(BootstrapVue)
export default {
  name: 'general',
  data () {
    return {
      currentPrivivka: null,
      itemsVACCINATIONS: [],
      selectedDate: null,
      selectedPreparat: null,
      selectedSeria: null,
      selectedDoza: null,
      nameUser: [],
      items: [],
      tabIndex: null,
      input_find: '',
      NumberKart: '',
      Fio: '',
      Pasport: '',
      DateRogd: '',
      Pol: '',
      selected: null,
      optionsPreparats: [
        {value: null, text: 'наименование препарата'},
        {value: 'АДСМ', text: 'АДСМ'}
      ],
      optionsSeria: [
        {value: null, text: 'серия препарата'},
        {value: 'П7', text: 'П7'},
        {value: 'П28', text: 'П28'}
      ],
      optionsDoza: [
        {value: null, text: 'доза'},
        {value: '0.5', text: '0.5'}
      ],
      fields: [
        {key: 'DATEPRIV', label: 'Дата прививки', class: 'text-justify col-xs-8'},
        {key: 'PREPARAT', label: 'Препарат', class: 'text-justify'},
        {key: 'SERIA', label: 'Серия', class: 'text-justify'},
        {key: 'DOZA', label: 'Доза', class: 'text-center'}]
    }
  },
  mounted: function () {
    this.CurrentUser()
  },
  methods: {
    logout: function () {
      Auth.logout()
    },
    setItemsDefaultVACCINATIONS (vaccination) {
      var ss = this
      controllers.setItemsDefaultVACCINATIONS(vaccination, this.NumberKart).then(function (response) {
        ss.itemsVACCINATIONS = JSON.parse(JSON.stringify(response.data))
        ss.currentPrivivka = vaccination
        alert(ss.currentPrivivka)
      }).catch(function (error) {
        console.log(error)
      })
    },
    CurrentUser: function () {
      var ss = this
      controllers.GetCurrentUser().then(function (response) {
        ss.nameUser = JSON.parse(JSON.stringify(response.data))
      }).catch(function (error) {
        console.log(error)
      })
    },
    addPrivivka: function () {
      var ss = this
      controllers.addPrivika(this.currentPrivivka,
        this.selectedDate,
        this.selectedPreparat,
        this.selectedSeria,
        this.selectedDoza,
        this.NumberKart
      ).then(function (response) {
        console.log(response.status)
        if (response.status === 200) {
          ss.setItemsDefaultVACCINATIONS('АДСМ')
        }
        if (response.status === 404) {

        }
      }).catch(function (error) {
        console.log(error.response)
        alert('Произошла ошибка, свяжитесь с системным администратором ' + error)
      })
    },
    startAutoUpdate: function () {
      this.timer = setInterval(this.GetModeVrach, 600000)
    },
    cancelAutoUpdate: function () {
      clearInterval(this.timer)
    },
    GetPatient () {
      localStorage.getItem('token_vaccinations')
      const data = new FormData()
      var ss = this
      data.append('patient', ss.input_find)
      axios.defaults.headers.common['Authorization'] = localStorage.getItem('token_vaccinations')
      axios.post('http://localhost:8084/jwt/FindPatientInArena', data, {responeType: 'application/json'})
        .then(function (response) {
          var data = JSON.parse(JSON.stringify(response.data))
          ss.NumberKart = data[0].NumberKart
          ss.Fio = data[0].Fio
          ss.Pasport = data[0].Pasport
          ss.DateRogd = data[0].DateRogd
          ss.Pol = data[0].Pol
          ss.setItemsDefaultVACCINATIONS('АДСМ')
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>
