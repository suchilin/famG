<template>
  <div>
    <v-row align="center" justify="center">
      <v-col cols="10">
        <v-card>
          <v-toolbar dense flat="true">
            <v-toolbar-title>Categorias</v-toolbar-title>

            <v-spacer></v-spacer>

            <v-btn icon>
              <v-icon @click="openAddCategoryDialog">mdi-plus</v-icon>
            </v-btn>
          </v-toolbar>
          <v-card-text>
            <v-container fluid grid-list-sm>
              <v-layout row wrap>
                <v-flex
                  v-for="category in categories"
                  :key="category.id"
                  xs12
                  sm6
                  md4
                  lg3
                >
                  <category-card />
                </v-flex>
              </v-layout>
            </v-container>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
    <category-form v-model="open" :title="title" />
  </div>
</template>
<script>
import CategoryCard from '@/components/CategoryCard.vue';
import CategoryForm from '@/components/CategoryForm.vue';
import { mapState, mapActions } from 'vuex';

export default {
  name: 'Categories',
  data: () => ({ open: false, title: '' }),
  components: {
    CategoryCard,
    CategoryForm,
  },
  computed: {
    ...mapState('categories', {
      categories: (state) => state.data,
      categoryError: (state) => state.error,
    }),
  },
  methods: {
    openAddCategoryDialog() {
      this.title = 'Add Category';
      this.open = true;
    },
    ...mapActions('categories', ['addCategory']),
  },
};
</script>
