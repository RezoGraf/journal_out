<template>
  <div id="btn-add-privika">
  <b-button variant="primary" size="sm" v-b-modal.ModalDifteria :disabled="this.input_find == ''">Добавить</b-button>
  <b-modal id="ModalDifteria"
           title="Добавить прививку от дефтерии "
           @ok="addPrivivka(this.currentPrivivka)">
    <b-container>
      <b-row>Дата вакцинации: <b-col align-self="center"><b-input v-model="selectedDate" type="date"></b-input></b-col></b-row> <br />
      <b-row>
        <b-form-select v-model="selectedPreparat" :options="optionsPreparats"></b-form-select>
      </b-row><br />
      <!--<b-row><b-input type="text" placeholder="Прививки"></b-input></b-row><br />-->
      <b-row>
        <b-form-select v-model="selectedSeria" :options="optionsSeria"></b-form-select>
      </b-row><br />
      <b-row>
        <b-form-select v-model="selectedDoza" :options="optionsDoza"></b-form-select>
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

  Vue.component('btn-add-privivka', {
    template: '#add-privivka',
    props: { selectedPrivika: String,
    }
  })
</script>
