<template>
  <header
    class="fixed inset-x-0 top-0 z-50 border-b border-gray-200/40 bg-white/70 px-6 py-4 backdrop-blur-xl dark:border-dark-800/50 dark:bg-dark-950/60"
  >
    <nav class="mx-auto flex max-w-6xl items-center justify-between">
      <router-link to="/home" class="flex items-center" :aria-label="siteName">
        <div class="h-10 w-10 overflow-hidden rounded-xl shadow-md">
          <img
            :src="siteLogo || '/logo.png'"
            alt="Logo"
            class="h-full w-full object-contain"
          />
        </div>
      </router-link>

      <div class="flex items-center gap-3">
        <LocaleSwitcher />

        <router-link
          v-for="item in navItems"
          :key="item.to"
          :to="item.to"
          class="rounded-lg p-2 text-gray-500 transition-colors hover:bg-gray-100 hover:text-gray-700 dark:text-dark-400 dark:hover:bg-dark-800 dark:hover:text-white"
          :title="item.title"
        >
          <Icon :name="item.icon" size="md" />
        </router-link>

        <button
          @click="$emit('toggle-theme')"
          class="rounded-lg p-2 text-gray-500 transition-colors hover:bg-gray-100 hover:text-gray-700 dark:text-dark-400 dark:hover:bg-dark-800 dark:hover:text-white"
          :title="isDark ? t('home.switchToLight') : t('home.switchToDark')"
        >
          <Icon v-if="isDark" name="sun" size="md" />
          <Icon v-else name="moon" size="md" />
        </button>

        <router-link
          v-if="isAuthenticated"
          :to="dashboardPath"
          class="inline-flex items-center gap-1.5 rounded-full bg-gray-900 py-1 pl-1 pr-2.5 transition-colors hover:bg-gray-800 dark:bg-gray-800 dark:hover:bg-gray-700"
        >
          <span
            class="flex h-5 w-5 items-center justify-center rounded-full bg-gradient-to-br from-primary-400 to-primary-600 text-[10px] font-semibold text-white"
          >
            {{ userInitial }}
          </span>
          <span class="text-xs font-medium text-white">{{
            t("home.dashboard")
          }}</span>
          <Icon
            name="externalLink"
            size="xs"
            class="text-gray-400"
            :stroke-width="2"
          />
        </router-link>
        <router-link
          v-else
          to="/login"
          class="inline-flex items-center rounded-full bg-gray-900 px-3 py-1 text-xs font-medium text-white transition-colors hover:bg-gray-800 dark:bg-gray-800 dark:hover:bg-gray-700"
        >
          {{ t("home.login") }}
        </router-link>
      </div>
    </nav>
  </header>
  <div class="h-[73px]" aria-hidden="true"></div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { useI18n } from "vue-i18n";
import LocaleSwitcher from "@/components/common/LocaleSwitcher.vue";
import Icon from "@/components/icons/Icon.vue";

type IconName = InstanceType<typeof Icon>["$props"]["name"];

defineProps<{
  siteName: string;
  siteLogo: string;
  isDark: boolean;
  isAuthenticated: boolean;
  dashboardPath: string;
  userInitial: string;
}>();

defineEmits<{
  "toggle-theme": [];
}>();

const { t } = useI18n();

const navItems = computed<Array<{ to: string; icon: IconName; title: string }>>(
  () => [
    {
      to: "/home",
      icon: "home",
      title: t("docs.backHome"),
    },
    {
      to: "/pricing",
      icon: "creditCard",
      title: t("nav.pricing"),
    },
    {
      to: "/faq",
      icon: "helpCircle",
      title: t("nav.faq"),
    },
    {
      to: "/docs",
      icon: "book",
      title: t("home.viewDocs"),
    },
  ],
);
</script>
