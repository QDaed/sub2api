<template>
  <div
    class="relative flex min-h-screen flex-col overflow-hidden bg-gradient-to-br from-gray-50 via-primary-50/30 to-gray-100 dark:from-dark-950 dark:via-dark-900 dark:to-dark-950"
  >
    <!-- Background Decorations -->
    <div class="pointer-events-none absolute inset-0 overflow-hidden">
      <div
        class="absolute -right-40 -top-40 h-96 w-96 rounded-full bg-primary-400/20 blur-3xl"
      ></div>
      <div
        class="absolute -bottom-40 -left-40 h-96 w-96 rounded-full bg-primary-500/15 blur-3xl"
      ></div>
      <div
        class="absolute inset-0 bg-[linear-gradient(rgba(20,184,166,0.03)_1px,transparent_1px),linear-gradient(90deg,rgba(20,184,166,0.03)_1px,transparent_1px)] bg-[size:64px_64px]"
      ></div>
    </div>

    <PublicNavbar
      :site-name="siteName"
      :site-logo="siteLogo"
      :is-dark="isDark"
      :is-authenticated="isAuthenticated"
      :dashboard-path="dashboardPath"
      :user-initial="userInitial"
      @toggle-theme="toggleTheme"
    />

    <main class="relative z-10 flex-1 px-6 py-16">
      <div class="mx-auto max-w-4xl">

        <!-- Header -->
        <div class="mb-12 text-center">
          <div
            class="mb-4 inline-flex items-center gap-2 rounded-full border border-primary-200 bg-white/70 px-4 py-2 text-sm font-medium text-primary-700 shadow-sm backdrop-blur-sm dark:border-primary-800 dark:bg-dark-800/70 dark:text-primary-300"
          >
            <Icon name="helpCircle" size="sm" />
            {{ t("faq.badge") }}
          </div>
          <h1 class="mb-4 text-4xl font-bold text-gray-900 dark:text-white md:text-5xl">
            {{ t("faq.title") }}
          </h1>
          <p class="text-lg text-gray-600 dark:text-dark-300">
            {{ t("faq.subtitle") }}
          </p>
        </div>

        <!-- FAQ Sections -->
        <div
          v-for="section in faqSections"
          :key="section.key"
          class="mb-10"
        >
          <h2 class="mb-4 text-xl font-bold text-gray-900 dark:text-white">
            {{ section.title }}
          </h2>
          <div class="space-y-3">
            <div
              v-for="item in section.items"
              :key="item.key"
              class="rounded-xl border border-gray-200/50 bg-white/60 backdrop-blur-sm dark:border-dark-700/50 dark:bg-dark-800/60"
            >
              <button
                @click="toggleFaq(section.key + '-' + item.key)"
                class="flex w-full items-center justify-between px-5 py-4 text-left"
              >
                <span class="font-medium text-gray-800 dark:text-dark-100">
                  {{ item.question }}
                </span>
                <Icon
                  :name="isOpen(section.key + '-' + item.key) ? 'chevronUp' : 'chevronDown'"
                  size="sm"
                  class="shrink-0 text-gray-400 transition-transform"
                />
              </button>
              <div
                v-if="isOpen(section.key + '-' + item.key)"
                class="border-t border-gray-100 px-5 py-4 text-sm leading-relaxed text-gray-600 dark:border-dark-700/50 dark:text-dark-400"
              >
                <div v-html="item.answer"></div>
              </div>
            </div>
          </div>
        </div>

        <!-- Still have questions -->
        <div
          class="rounded-2xl border border-primary-200/50 bg-gradient-to-r from-primary-50/80 to-blue-50/80 p-8 text-center dark:border-primary-800/30 dark:from-primary-900/20 dark:to-blue-900/20"
        >
          <h3 class="mb-2 text-xl font-bold text-gray-900 dark:text-white">
            {{ t("faq.stillHaveQuestions.title") }}
          </h3>
          <p class="mb-4 text-gray-600 dark:text-dark-300">
            {{ t("faq.stillHaveQuestions.description") }}
          </p>
          <div class="flex flex-wrap items-center justify-center gap-3">
            <a
              v-if="telegramSupport"
              :href="telegramSupport"
              target="_blank"
              rel="noopener noreferrer"
              class="inline-flex items-center gap-2 rounded-full bg-primary-500 px-5 py-2.5 text-sm font-medium text-white shadow transition-colors hover:bg-primary-600"
            >
              <svg class="h-4 w-4" fill="currentColor" viewBox="0 0 24 24">
                <path d="M11.944 0A12 12 0 0 0 0 12a12 12 0 0 0 12 12 12 12 0 0 0 12-12A12 12 0 0 0 12 0a12 12 0 0 0-.056 0zm4.962 7.224c.1-.002.321.023.465.14a.506.506 0 0 1 .171.325c.016.093.036.306.02.472-.18 1.898-.962 6.502-1.36 8.627-.168.9-.499 1.201-.82 1.23-.696.065-1.225-.46-1.9-.902-1.056-.693-1.653-1.124-2.678-1.8-1.185-.78-.417-1.21.258-1.91.177-.184 3.247-2.977 3.307-3.23.007-.032.014-.15-.056-.212s-.174-.041-.249-.024c-.106.024-1.793 1.14-5.061 3.345-.48.33-.913.49-1.302.48-.428-.008-1.252-.241-1.865-.44-.752-.245-1.349-.374-1.297-.789.027-.216.325-.437.893-.663 3.498-1.524 5.83-2.529 6.998-3.014 3.332-1.386 4.025-1.627 4.476-1.635z"/>
              </svg>
              {{ t("faq.stillHaveQuestions.telegram") }}
            </a>
            <router-link
              to="/register"
              class="inline-flex items-center gap-2 rounded-full bg-gray-900 px-5 py-2.5 text-sm font-medium text-white shadow transition-colors hover:bg-gray-800 dark:bg-gray-800 dark:hover:bg-gray-700"
            >
              {{ t("faq.stillHaveQuestions.getStarted") }}
              <Icon name="arrowRight" size="sm" />
            </router-link>
          </div>
        </div>

      </div>
    </main>

    <!-- Footer -->
    <footer
      class="relative z-10 border-t border-gray-200/50 px-6 py-8 dark:border-dark-800/50"
    >
      <div
        class="mx-auto flex max-w-4xl flex-col items-center justify-center gap-4 text-center sm:flex-row sm:text-left"
      >
        <p class="text-sm text-gray-500 dark:text-dark-400">
          &copy; {{ currentYear }} {{ siteName }}.
          {{ t("home.footer.allRightsReserved") }}
        </p>
        <div class="flex items-center gap-4">
          <router-link
            to="/docs"
            class="text-sm text-gray-500 transition-colors hover:text-gray-700 dark:text-dark-400 dark:hover:text-white"
          >
            {{ t("home.docs") }}
          </router-link>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { useI18n } from "vue-i18n";
import { useAuthStore, useAppStore } from "@/stores";
import Icon from "@/components/icons/Icon.vue";
import PublicNavbar from "@/components/layout/PublicNavbar.vue";

const { t } = useI18n();

const authStore = useAuthStore();
const appStore = useAppStore();

const siteName = computed(
  () =>
    appStore.cachedPublicSettings?.site_name || appStore.siteName || "AISHOPACC",
);
const siteLogo = computed(
  () => appStore.cachedPublicSettings?.site_logo || appStore.siteLogo || "",
);

const telegramSupport = computed(() =>
  (appStore.cachedPublicSettings as Record<string, unknown>)?.support_telegram as string ||
  "https://t.me/shopaikey_support",
);

const isAuthenticated = computed(() => authStore.isAuthenticated);
const isAdmin = computed(() => authStore.isAdmin);
const dashboardPath = computed(() =>
  isAdmin.value ? "/admin/dashboard" : "/dashboard",
);
const userInitial = computed(() => {
  const user = authStore.user;
  if (!user || !user.email) return "";
  return user.email.charAt(0).toUpperCase();
});

const currentYear = computed(() => new Date().getFullYear());

// Theme
const isDark = ref(document.documentElement.classList.contains("dark"));

function toggleTheme() {
  isDark.value = !isDark.value;
  document.documentElement.classList.toggle("dark", isDark.value);
  localStorage.setItem("theme", isDark.value ? "dark" : "light");
}

function initTheme() {
  const savedTheme = localStorage.getItem("theme");
  if (
    savedTheme === "dark" ||
    (!savedTheme && window.matchMedia("(prefers-color-scheme: dark)").matches)
  ) {
    isDark.value = true;
    document.documentElement.classList.add("dark");
  }
}

// FAQ accordion state
const openFaqs = ref<Set<string>>(new Set());

function toggleFaq(key: string) {
  if (openFaqs.value.has(key)) {
    openFaqs.value.delete(key);
  } else {
    openFaqs.value.add(key);
  }
}

function isOpen(key: string) {
  return openFaqs.value.has(key);
}

// FAQ content
const faqSections = computed(() => [
  {
    key: "general",
    title: t("faq.sections.general.title"),
    items: [
      {
        key: "what",
        question: t("faq.sections.general.what.question"),
        answer: t("faq.sections.general.what.answer"),
      },
      {
        key: "why",
        question: t("faq.sections.general.why.question"),
        answer: t("faq.sections.general.why.answer"),
      },
      {
        key: "security",
        question: t("faq.sections.general.security.question"),
        answer: t("faq.sections.general.security.answer"),
      },
    ],
  },
  {
    key: "payment",
    title: t("faq.sections.payment.title"),
    items: [
      {
        key: "pricing",
        question: t("faq.sections.payment.pricing.question"),
        answer: t("faq.sections.payment.pricing.answer"),
      },
      {
        key: "expiry",
        question: t("faq.sections.payment.expiry.question"),
        answer: t("faq.sections.payment.expiry.answer"),
      },
      {
        key: "refund",
        question: t("faq.sections.payment.refund.question"),
        answer: t("faq.sections.payment.refund.answer"),
      },
    ],
  },
  {
    key: "usage",
    title: t("faq.sections.usage.title"),
    items: [
      {
        key: "multi-device",
        question: t("faq.sections.usage.multiDevice.question"),
        answer: t("faq.sections.usage.multiDevice.answer"),
      },
      {
        key: "balance",
        question: t("faq.sections.usage.balance.question"),
        answer: t("faq.sections.usage.balance.answer"),
      },
      {
        key: "rate-limit",
        question: t("faq.sections.usage.rateLimit.question"),
        answer: t("faq.sections.usage.rateLimit.answer"),
      },
    ],
  },
  {
    key: "integration",
    title: t("faq.sections.integration.title"),
    items: [
      {
        key: "sdk",
        question: t("faq.sections.integration.sdk.question"),
        answer: t("faq.sections.integration.sdk.answer"),
      },
      {
        key: "compatible",
        question: t("faq.sections.integration.compatible.question"),
        answer: t("faq.sections.integration.compatible.answer"),
      },
      {
        key: "tools",
        question: t("faq.sections.integration.tools.question"),
        answer: t("faq.sections.integration.tools.answer"),
      },
    ],
  },
  {
    key: "support",
    title: t("faq.sections.support.title"),
    items: [
      {
        key: "contact",
        question: t("faq.sections.support.contact.question"),
        answer: t("faq.sections.support.contact.answer"),
      },
      {
        key: "bulk",
        question: t("faq.sections.support.bulk.question"),
        answer: t("faq.sections.support.bulk.answer"),
      },
    ],
  },
]);

onMounted(() => {
  initTheme();
  authStore.checkAuth();
  if (!appStore.publicSettingsLoaded) {
    appStore.fetchPublicSettings();
  }
});
</script>
