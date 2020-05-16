<template>
  <v-card class="ma-1">
    <v-img
      src="https://cdn.vuetifyjs.com/images/cards/sunshine.jpg"
      height="200px"
    ></v-img>

    <v-card-title>{{ title }}</v-card-title>

    <v-card-subtitle>
      {{ description }}
    </v-card-subtitle>

    <v-card-actions>
      <v-btn text>Ver productos ...</v-btn>
      <v-spacer></v-spacer>

      <v-btn icon>
        <v-icon>edit</v-icon>
      </v-btn>
      <v-btn icon>
        <v-icon>delete</v-icon>
      </v-btn>
    </v-card-actions>

    <v-expand-transition>
      <div v-show="show">
        <v-divider></v-divider>

        <v-card-text> </v-card-text>
      </div>
    </v-expand-transition>
  </v-card> </template
>;

<script>
export default {
  props: ['title', 'description', 'image'],
};
</script>
