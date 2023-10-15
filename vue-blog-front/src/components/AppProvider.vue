<template>
    <n-config-provider :theme-overrides="themeOverrides">
        <slot></slot>
    </n-config-provider>
</template>
  
<script setup>
const themeOverrides = {
    common: {
        primaryColor: '#316C72FF',
        primaryColorHover: '#316C72E3',
        primaryColorPressed: '#2B4C59FF',
        primaryColorSuppl: '#316C7263',
    },
}
</script>
  
  