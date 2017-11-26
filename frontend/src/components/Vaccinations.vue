<template>
  <div class="general">
      <b-container>
      <br>
        <b-row class="mb-4 text-right">
          <b-col cols="3">
            <div class="input-group">
            <b-form-input :formatter="format" @keyup.native.delete="clearSearchPatient" v-model="input_find" placeholder="№ Карты"></b-form-input>
              <b-button @click="clearSearchPatient" class="input-group-addon">x</b-button>
            </div>
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
                <b-tab title="Дифтерия" @click="setItemsDefaultVACCINATIONS('Дефтерия')">
                  <b-table striped hover
                           :items="itemsVACCINATIONS"
                           :fields="fields">
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

                  <b-button variant="primary" size="sm" v-b-modal.ModalAddVaccinationDEFTERIA :disabled="this.input_find == ''">Добавить</b-button>

                </b-tab>
                  <b-tab title="Клещевой энцефалит" @click="setItemsDefaultVACCINATIONS('Клещевой энцефалит')">
                    <b-table striped hover
                             :items="itemsVACCINATIONS"
                             :fields="fields">
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

                    <b-button variant="primary" size="sm" v-b-modal.ModalAddVaccinationKLESH :disabled="this.input_find == ''">Добавить</b-button>

                </b-tab>
                <b-tab title="Корь" @click="setItemsDefaultVACCINATIONS('Корь')">
                  <b-table striped hover
                           :items="itemsVACCINATIONS"
                           :fields="fields">
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

                  <b-button variant="primary" size="sm" v-b-modal.ModalAddVaccinationKOR :disabled="this.input_find == ''">Добавить</b-button>

                </b-tab>
                <b-tab title="Краснуха" @click="setItemsDefaultVACCINATIONS('Краснуха')">
                  <b-table striped hover
                           :items="itemsVACCINATIONS"
                           :fields="fields">
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

                  <b-button variant="primary" size="sm" v-b-modal.ModalAddVaccinationKRASNUHA :disabled="this.input_find == ''">Добавить</b-button>

                </b-tab>
                <b-tab title="Гепатит В" @click="setItemsDefaultVACCINATIONS('Гепатит В')">
                  <b-table striped hover
                           :items="itemsVACCINATIONS"
                           :fields="fields">
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

                  <b-button variant="primary" size="sm" v-b-modal.ModalAddVaccinationGEPATITB :disabled="this.input_find == ''">Добавить</b-button>

                </b-tab>
                <b-tab title="Грипп" @click="setItemsDefaultVACCINATIONS('Грипп')">
                  <b-table striped hover
                           :items="itemsVACCINATIONS"
                           :fields="fields">
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
                  <b-button variant="primary" size="sm" v-b-modal.ModalAddVaccinationGRIPP :disabled="this.input_find == ''">Добавить</b-button>
                </b-tab>
              </b-tabs>
            </b-card>
          </b-col>

        </b-row>
      </b-container>

    <b-modal id="ModalAddVaccinationDEFTERIA"
                           title = "Новая прививка Дефтерия"
                           @ok="addPrivivka()">
    <b-container>
      <b-row>Дата вакцинации: <b-col align-self="center"><b-input v-model="selectedDateDEFTERIA" type="date"></b-input></b-col></b-row> <br />
      <b-row>
        <b-form-select v-model="selectedPreparatDEFTERIA" :options="optionsPreparatsDEFTERIA"></b-form-select>
      </b-row><br />
      <!--<b-row><b-input type="text" placeholder="Прививки"></b-input></b-row><br />-->
      <b-row>
        <b-form-select v-model="selectedSeriaDEFTERIA" :options="optionsSeriaDEFTERIA"></b-form-select>
      </b-row><br />
      <b-row>
        <b-form-select v-model="selectedDozaDEFTERIA" :options="optionsDozaDEFTERIA"></b-form-select>
      </b-row><br />
      <!--<b-row><b-input type="text" placeholder="Медотвод"></b-input></b-row>-->
    </b-container>
  </b-modal>
    <b-modal id="ModalAddVaccinationKLESH"
             title = "Новая прививка Клещевой энцефалит"
             @ok="addPrivivka()">
      <b-container>
        <b-row>Дата вакцинации: <b-col align-self="center"><b-input v-model="selectedDateKLESH" type="date"></b-input></b-col></b-row> <br />
        <b-row>
          <b-form-input v-model="selectedPreparatKLESH" placeholder="Препарат"></b-form-input>
        </b-row><br />
        <!--<b-row><b-input type="text" placeholder="Прививки"></b-input></b-row><br />-->
        <b-row>
          <b-form-input v-model="selectedSeriaKLESH" placeholder="Серия"></b-form-input>
        </b-row><br />
        <b-row>
          <b-form-select v-model="selectedDozaKLESH" :options="optionsDozaKLESH"></b-form-select>
        </b-row><br />
        <!--<b-row><b-input type="text" placeholder="Медотвод"></b-input></b-row>-->
      </b-container>
    </b-modal>
    <b-modal id="ModalAddVaccinationKOR"
             title = "Новая прививка Корь"
             @ok="addPrivivka()">
      <b-container>
        <b-row>Дата вакцинации: <b-col align-self="center"><b-input v-model="selectedDateKOR" type="date"></b-input></b-col></b-row> <br />
        <b-row>
          <b-form-select v-model="selectedPreparatKOR" :options="optionsPreparatsKOR"></b-form-select>
        </b-row><br />
        <!--<b-row><b-input type="text" placeholder="Прививки"></b-input></b-row><br />-->
        <b-row>
          <b-form-select v-model="selectedSeriaKOR" :options="optionsSeriaKOR"></b-form-select>
        </b-row><br />
        <b-row>
          <b-form-select v-model="selectedDozaKOR" :options="optionsDozaKOR"></b-form-select>
        </b-row><br />
        <!--<b-row><b-input type="text" placeholder="Медотвод"></b-input></b-row>-->
      </b-container>
    </b-modal>
    <b-modal id="ModalAddVaccinationKRASNUHA"
             title = "Новая прививка Краснуха"
             @ok="addPrivivka()">
      <b-container>
        <b-row>Дата вакцинации: <b-col align-self="center"><b-input v-model="selectedDateKRASNUHA" type="date"></b-input></b-col></b-row> <br />
        <b-row>
          <b-form-input v-model="selectedPreparatKRASNUHA" placeholder="Препарат"></b-form-input>
        </b-row><br />
        <!--<b-row><b-input type="text" placeholder="Прививки"></b-input></b-row><br />-->
        <b-row>
          <b-form-input v-model="selectedSeriaKRASNUHA" placeholder="Серия"></b-form-input>
        </b-row><br />
        <b-row>
          <b-form-select v-model="selectedDozaKRASNUHA" :options="optionsDozaKLESH"></b-form-select>
        </b-row><br />
        <!--<b-row><b-input type="text" placeholder="Медотвод"></b-input></b-row>-->
      </b-container>
    </b-modal>
    <b-modal id="ModalAddVaccinationGEPATITB"
             title = "Новая прививка Гепатит В"
             @ok="addPrivivka()">
      <b-container>
        <b-row>Дата вакцинации: <b-col align-self="center"><b-input v-model="selectedDateGEPATITB" type="date"></b-input></b-col></b-row> <br />
        <b-row>
          <b-form-select v-model="selectedPreparatGEPATITB" :options="optionsPreparatsGEPATITB"></b-form-select>
        </b-row><br />
        <!--<b-row><b-input type="text" placeholder="Прививки"></b-input></b-row><br />-->
        <b-row>
          <b-form-select v-model="selectedSeriaGEPATITB" :options="optionsSeriaGEPATITB"></b-form-select>
        </b-row><br />
        <b-row>
          <b-form-select v-model="selectedDozaGEPATITB" :options="optionsDozaGEPATITB"></b-form-select>
        </b-row><br />
        <!--<b-row><b-input type="text" placeholder="Медотвод"></b-input></b-row>-->
      </b-container>
    </b-modal>
    <b-modal id="ModalAddVaccinationGRIPP"
             title = "Новая прививка Грипп"
             @ok="addPrivivka()">
      <b-container>
        <b-row>Дата вакцинации: <b-col align-self="center"><b-input v-model="selectedDateGRIPP" type="date"></b-input></b-col></b-row> <br />
        <b-row>
          <b-form-select v-model="selectedPreparatGRIPP" :options="optionsPreparatsGRIPP"></b-form-select>
        </b-row><br />
        <!--<b-row><b-input type="text" placeholder="Прививки"></b-input></b-row><br />-->
        <b-row>
          <b-form-input v-model="selectedSeriaGRIPP" placeholder="Серия"></b-form-input>
        </b-row><br />
        <b-row>
          <b-form-select v-model="selectedDozaGRIPP" :options="optionsDozaGRIPP"></b-form-select>
        </b-row><br />
        <!--<b-row><b-input type="text" placeholder="Медотвод"></b-input></b-row>-->
      </b-container>
    </b-modal>

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
      flag: false,
      currentPrivivka: null,
      itemsVACCINATIONS: [],
      selectedDateDEFTERIA: null,
      selectedPreparatDEFTERIA: null,
      selectedSeriaDEFTERIA: null,
      selectedDozaDEFTERIA: null,
      selectedDateKLESH: null,
      selectedPreparatKLESH: null,
      selectedSeriaKLESH: null,
      selectedDozaKLESH: null,
      selectedDateKOR: null,
      selectedPreparatKOR: null,
      selectedSeriaKOR: null,
      selectedDozaKOR: null,
      selectedDateKRASNUHA: null,
      selectedPreparatKRASNUHA: null,
      selectedSeriaKRASNUHA: null,
      selectedDozaKRASNUHA: null,
      selectedDateGEPATITB: null,
      selectedPreparatGEPATITB: null,
      selectedSeriaGEPATITB: null,
      selectedDozaGEPATITB: null,
      selectedDateGRIPP: null,
      selectedPreparatGRIPP: null,
      selectedSeriaGRIPP: null,
      selectedDozaGRIPP: null,
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
      optionsPreparatsDEFTERIA: [
        {value: null, text: 'наименование препарата'},
        {value: 'АДСМ', text: 'АДСМ'}
      ],
      optionsSeriaDEFTERIA: [
        {value: null, text: 'серия препарата'},
        {value: 'П7', text: 'П7'},
        {value: 'П28', text: 'П28'}
      ],
      optionsDozaDEFTERIA: [
        {value: null, text: 'доза'},
        {value: '0.5', text: '0.5'}
      ],
      optionsDozaKLESH: [
        {value: null, text: 'доза'},
        {value: '0.5', text: '0.5'}
      ],
      optionsPreparatsKOR: [
        {value: null, text: 'наименование препарата'},
        {value: 'ЖКВ', text: 'ЖКВ'}
      ],
      optionsSeriaKOR: [
        {value: null, text: 'серия препарата'},
        {value: 'M200', text: 'M200'}
      ],
      optionsDozaKOR: [
        {value: null, text: 'доза'},
        {value: '0.5', text: '0.5'}
      ],
      optionsDozaKRASNUHA: [
        {value: null, text: 'доза'},
        {value: '0.5', text: '0.5'}
      ],
      optionsPreparatsGEPATITB: [
        {value: null, text: 'наименование препарата'},
        {value: 'Регевак', text: 'Регевак'}
      ],
      optionsSeriaGEPATITB: [
        {value: null, text: 'серия препарата'},
        {value: '130216', text: '130216'}
      ],
      optionsDozaGEPATITB: [
        {value: null, text: 'доза'},
        {value: '1.0', text: '1.0'}
      ],
      optionsPreparatsGRIPP: [
        {value: null, text: 'наименование препарата'},
        {value: 'Совигрипп', text: 'Совигрипп'}
      ],
      optionsDozaGRIPP: [
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
    format: function (value, event) {
      const val = value.replace(/(?!-)[^0-9.]/g, ' ')
      if (this.flag) {
        this.clearSearchPatient()
        this.flag = false
      }
      return val
    },
    logout: function () {
      Auth.logout()
    },
    clearSearchPatient: function () {
      this.input_find = ''
      this.Fio = ''
      this.Pasport = ''
      this.DateRogd = ''
      this.NumberKart = ''
      this.Pol = ''
      this.itemsVACCINATIONS = null
    },
    setItemsDefaultVACCINATIONS (vaccination) {
      if (this.NumberKart !== '') {
        var ss = this
        controllers.setItemsDefaultVACCINATIONS(vaccination, this.NumberKart).then(function (response) {
          ss.itemsVACCINATIONS = JSON.parse(JSON.stringify(response.data))
          ss.currentPrivivka = vaccination
        }).catch(function (error) {
          console.log(error)
        })
      }
    },
    CurrentUser: function () {
      var ss = this
      controllers.GetCurrentUser().then(function (response) {
        ss.nameUser = JSON.parse(JSON.stringify(response.data))
      }).catch(function (error) {
        console.log(error)
      })
    },
    GetPatient () {
      this.flag = true
      localStorage.getItem('token_vaccinations')
      const data = new FormData()
      var ss = this
      data.append('patient', ss.input_find.trim())
      axios.defaults.headers.common['Authorization'] = localStorage.getItem('token_vaccinations')
      axios.post('http://localhost:8084/jwt/FindPatientInArena', data, {responeType: 'application/json'})
        .then(function (response) {
          var data = JSON.parse(JSON.stringify(response.data))
          ss.NumberKart = data[0].NumberKart
          ss.Fio = data[0].Fio
          ss.Pasport = data[0].Pasport
          ss.DateRogd = data[0].DateRogd
          ss.Pol = data[0].Pol
          ss.setItemsDefaultVACCINATIONS('Дефтерия')
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    addPrivivka: function () {
      var ss = this
      var selectedDate, selectedPreparat, selectedSeria, selectedDoza
      if (this.currentPrivivka === 'Дефтерия') {
        selectedDate = ss.selectedDateDEFTERIA
        selectedPreparat = ss.selectedPreparatDEFTERIA
        selectedSeria = ss.selectedSeriaDEFTERIA
        selectedDoza = ss.selectedDozaDEFTERIA
      }
      if (this.currentPrivivka === 'Клещевой энцефалит') {
        selectedDate = ss.selectedDateKLESH
        selectedPreparat = ss.selectedPreparatKLESH
        selectedSeria = ss.selectedSeriaKLESH
        selectedDoza = ss.selectedDozaKLESH
      }
      if (this.currentPrivivka === 'Корь') {
        selectedDate = ss.selectedDateKOR
        selectedPreparat = ss.selectedPreparatKOR
        selectedSeria = ss.selectedSeriaKOR
        selectedDoza = ss.selectedDozaKOR
      }
      if (this.currentPrivivka === 'Краснуха') {
        selectedDate = ss.selectedDateKRASNUHA
        selectedPreparat = ss.selectedPreparatKRASNUHA
        selectedSeria = ss.selectedSeriaKRASNUHA
        selectedDoza = ss.selectedDozaKRASNUHA
      }
      if (this.currentPrivivka === 'Гепатит В') {
        selectedDate = ss.selectedDateGEPATITB
        selectedPreparat = ss.selectedPreparatGEPATITB
        selectedSeria = ss.selectedSeriaGEPATITB
        selectedDoza = ss.selectedDozaGEPATITB
      }
      if (this.currentPrivivka === 'Грипп') {
        selectedDate = ss.selectedDateGRIPP
        selectedPreparat = ss.selectedPreparatGRIPP
        selectedSeria = ss.selectedSeriaGRIPP
        selectedDoza = ss.selectedDozaGRIPP
      }
      controllers.addPrivika(this.currentPrivivka,
        selectedDate,
        selectedPreparat,
        selectedSeria,
        selectedDoza,
        this.NumberKart
      ).then(function (response) {
        console.log(response.status)
        if (response.status === 200) {
          ss.setItemsDefaultVACCINATIONS(ss.currentPrivivka)
          ss.clearSelected()
        }
        if (response.status === 404) {

        }
      }).catch(function (error) {
        console.log(error.response)
        alert('Произошла ошибка, свяжитесь с системным администратором ' + error)
      })
    },
    clearSelected: function () {
      this.selectedDateDEFTERIA = null
      this.selectedPreparatDEFTERIA = null
      this.selectedSeriaDEFTERIA = null
      this.selectedDozaDEFTERIA = null
      this.selectedDateKLESH = null
      this.selectedPreparatKLESH = null
      this.selectedSeriaKLESH = null
      this.selectedDozaKLESH = null
      this.selectedDateKOR = null
      this.selectedPreparatKOR = null
      this.selectedSeriaKOR = null
      this.selectedDozaKOR = null
      this.selectedDateKRASNUHA = null
      this.selectedPreparatKRASNUHA = null
      this.selectedSeriaKRASNUHA = null
      this.selectedDozaKRASNUHA = null
      this.selectedDateGEPATITB = null
      this.selectedPreparatGEPATITB = null
      this.selectedSeriaGEPATITB = null
      this.selectedDozaGEPATITB = null
      this.selectedDateGRIPP = null
      this.selectedPreparatGRIPP = null
      this.selectedSeriaGRIPP = null
      this.selectedDozaGRIPP = null
    }
  }
}
</script>
