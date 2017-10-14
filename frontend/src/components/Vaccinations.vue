<template>
  <div class="general">
  <div class="container">
    <div class="row">
      <div class="col">
        <b-button v-b-modal.modallg variant="primary">Добавить пациента</b-button>
      </div>
      </div>
    <div class="row">
  <div class="col">

  </div>
    </div>
    </div>
    <b-modal id="modallg" size="lg" title="Добавить пациента">
      <b-container fluid>
        <b-row class="mb-1 text-center">
          <b-col cols="3">
            <b-form-input placeholder="№ Карты"></b-form-input>
          </b-col>
          <b-col cols="8">
            <b-form-input placeholder="Ф.И.О"></b-form-input>
            </b-col>
        </b-row>
        <b-row class="mb1 text-center">
            <b-col cols="4">
              <b-form-input placeholder="Паспорт"></b-form-input>
            </b-col>
          <b-col cols="5">
            <b-form-select v-model="selected" :options="options" class="mb-3">
            </b-form-select>
          </b-col>
            <b-col cols="3">
              <b-form-input type="date"></b-form-input>
            </b-col>
        </b-row>
        <b-row class="mb1 text-center">
          <b-col cols="8">

          </b-col>
        </b-row>
      </b-container>
    </b-modal>
  </div>
</template>

<script>
import Vue from 'vue'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.use(BootstrapVue)

export default {
  name: 'general',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App',
      selected: null,
      options: [
        {value: null, text: 'Кем выдан'},
        {value: 'a', text: 'УВД Ленинского района'},
        {value: 'b', text: 'УВД Кировского района'},
        {value: 'c', text: 'УВД Центрального района'}
      ]
    }
  }
}
</script>
