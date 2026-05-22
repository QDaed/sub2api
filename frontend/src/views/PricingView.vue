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
      <div class="mx-auto max-w-6xl">

        <!-- Header -->
        <div class="mb-12 text-center">
          <div
            class="mb-4 inline-flex items-center gap-2 rounded-full border border-primary-200 bg-white/70 px-4 py-2 text-sm font-medium text-primary-700 shadow-sm backdrop-blur-sm dark:border-primary-800 dark:bg-dark-800/70 dark:text-primary-300"
          >
            <Icon name="creditCard" size="sm" />
            {{ t("pricing.badge") }}
          </div>
          <h1 class="mb-4 text-4xl font-bold text-gray-900 dark:text-white md:text-5xl">
            {{ t("pricing.title") }}
          </h1>
          <p class="text-lg text-gray-600 dark:text-dark-300">
            {{ t("pricing.subtitle") }}
          </p>
        </div>

        <!-- Pricing Cards -->
        <div class="mb-16 grid gap-6 md:grid-cols-2 lg:grid-cols-3">
          <div
            v-for="plan in plans"
            :key="plan.id"
            class="group relative rounded-2xl border p-6 transition-all duration-300"
            :class="[
              plan.popular
                ? 'border-primary-400 bg-white/80 shadow-xl shadow-primary-500/10 dark:border-primary-600 dark:bg-dark-800/80'
                : 'border-gray-200/50 bg-white/60 backdrop-blur-sm dark:border-dark-700/50 dark:bg-dark-800/60 hover:shadow-lg hover:shadow-primary-500/5',
            ]"
          >
            <div
              v-if="plan.popular"
              class="absolute -top-3 left-1/2 -translate-x-1/2 rounded-full bg-primary-500 px-4 py-1 text-xs font-semibold text-white shadow"
            >
              {{ t("pricing.popular") }}
            </div>
            <div class="mb-4">
              <h3 class="text-lg font-bold text-gray-900 dark:text-white">
                {{ plan.name }}
              </h3>
              <p class="mt-1 text-sm text-gray-500 dark:text-dark-400">
                {{ plan.description }}
              </p>
            </div>
            <div class="mb-6">
              <span class="text-3xl font-bold text-gray-900 dark:text-white">
                {{ plan.priceLabel }}
              </span>
              <span
                v-if="plan.priceNote"
                class="text-sm text-gray-500 dark:text-dark-400"
              >
                {{ plan.priceNote }}
              </span>
            </div>
            <ul class="mb-6 space-y-2">
              <li
                v-for="feature in plan.features"
                :key="feature"
                class="flex items-start gap-2 text-sm text-gray-600 dark:text-dark-400"
              >
                <Icon
                  name="check"
                  size="sm"
                  class="mt-0.5 shrink-0 text-primary-500"
                />
                {{ feature }}
              </li>
            </ul>
            <router-link
              :to="isAuthenticated ? '/purchase' : '/register'"
              class="block w-full rounded-xl py-2.5 text-center text-sm font-semibold transition-colors"
              :class="plan.popular
                ? 'bg-primary-500 text-white shadow shadow-primary-500/30 hover:bg-primary-600'
                : 'bg-gray-900 text-white hover:bg-gray-800 dark:bg-gray-800 dark:hover:bg-gray-700'"
            >
              {{ plan.buttonText }}
            </router-link>
          </div>
        </div>

        <!-- Pay-per-use Section -->
        <div class="mb-16 rounded-2xl border border-gray-200/50 bg-white/60 p-8 backdrop-blur-sm dark:border-dark-700/50 dark:bg-dark-800/60">
          <div class="mb-6 text-center">
            <h2 class="text-2xl font-bold text-gray-900 dark:text-white">
              {{ t("pricing.payPerUse.title") }}
            </h2>
            <p class="mt-2 text-gray-600 dark:text-dark-300">
              {{ t("pricing.payPerUse.subtitle") }}
            </p>
          </div>
          <div class="mb-6 overflow-x-auto">
            <table class="w-full min-w-[600px] text-sm">
              <thead>
                <tr class="border-b border-gray-200/50 dark:border-dark-700/50">
                  <th class="px-4 py-3 text-left font-semibold text-gray-600 dark:text-dark-300">
                    {{ t("pricing.table.model") }}
                  </th>
                  <th class="px-4 py-3 text-right font-semibold text-gray-600 dark:text-dark-300">
                    {{ t("pricing.table.input") }}
                  </th>
                  <th class="px-4 py-3 text-right font-semibold text-gray-600 dark:text-dark-300">
                    {{ t("pricing.table.output") }}
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="group in modelGroups"
                  :key="group.category"
                  class="border-b border-gray-100 dark:border-dark-700/30 last:border-0"
                >
                  <td
                    colspan="3"
                    class="px-4 py-2 text-xs font-semibold uppercase tracking-wider text-primary-600 dark:text-primary-400"
                  >
                    {{ group.category }}
                  </td>
                </tr>
                <tr
                  v-for="model in flatModels"
                  :key="model.name"
                  class="border-b border-gray-100 dark:border-dark-700/30 last:border-0"
                >
                  <td class="px-4 py-2.5 font-medium text-gray-800 dark:text-dark-100">
                    {{ model.name }}
                  </td>
                  <td class="px-4 py-2.5 text-right text-gray-600 dark:text-dark-400">
                    {{ model.input }}
                  </td>
                  <td class="px-4 py-2.5 text-right text-gray-600 dark:text-dark-400">
                    {{ model.output }}
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
          <div class="rounded-xl border border-blue-100 bg-blue-50/60 p-4 text-sm text-blue-700 dark:border-blue-900/30 dark:bg-blue-900/10 dark:text-blue-300">
            <div class="flex items-start gap-2">
              <Icon name="info" size="sm" class="mt-0.5 shrink-0" />
              <span>{{ t("pricing.table.note") }}</span>
            </div>
          </div>
        </div>

        <!-- CTA -->
        <div
          class="rounded-2xl border border-primary-200/50 bg-gradient-to-r from-primary-50/80 to-blue-50/80 p-8 text-center dark:border-primary-800/30 dark:from-primary-900/20 dark:to-blue-900/20 md:p-12"
        >
          <h2 class="mb-3 text-2xl font-bold text-gray-900 dark:text-white md:text-3xl">
            {{ t("pricing.cta.title") }}
          </h2>
          <p class="mb-6 text-lg text-gray-600 dark:text-dark-300">
            {{ t("pricing.cta.description") }}
          </p>
          <div class="flex flex-wrap items-center justify-center gap-3">
            <router-link
              :to="isAuthenticated ? '/dashboard' : '/register'"
              class="btn btn-primary px-8 py-3 shadow-lg shadow-primary-500/30"
            >
              {{ isAuthenticated ? t("pricing.cta.dashboard") : t("pricing.cta.signup") }}
              <Icon name="arrowRight" size="md" class="ml-2" />
            </router-link>
            <router-link
              to="/faq"
              class="inline-flex items-center gap-2 rounded-full bg-gray-900/10 px-6 py-3 text-sm font-medium text-gray-700 transition-colors hover:bg-gray-900/20 dark:bg-white/10 dark:text-dark-200 dark:hover:bg-white/20"
            >
              {{ t("pricing.cta.faq") }}
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
        class="mx-auto flex max-w-6xl flex-col items-center justify-center gap-4 text-center sm:flex-row sm:text-left"
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
          <router-link
            to="/faq"
            class="text-sm text-gray-500 transition-colors hover:text-gray-700 dark:text-dark-400 dark:hover:text-white"
          >
            {{ t("faq.badge") }}
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

// Plans
const plans = computed(() => [
  {
    id: "free",
    name: t("pricing.plans.free.name"),
    description: t("pricing.plans.free.description"),
    priceLabel: t("pricing.plans.free.priceLabel"),
    priceNote: "",
    popular: false,
    features: [
      t("pricing.plans.free.f1"),
      t("pricing.plans.free.f2"),
      t("pricing.plans.free.f3"),
    ],
    buttonText: t("pricing.plans.free.button"),
  },
  {
    id: "starter",
    name: t("pricing.plans.starter.name"),
    description: t("pricing.plans.starter.description"),
    priceLabel: t("pricing.plans.starter.priceLabel"),
    priceNote: t("pricing.plans.starter.priceNote"),
    popular: true,
    features: [
      t("pricing.plans.starter.f1"),
      t("pricing.plans.starter.f2"),
      t("pricing.plans.starter.f3"),
      t("pricing.plans.starter.f4"),
    ],
    buttonText: t("pricing.plans.starter.button"),
  },
  {
    id: "pro",
    name: t("pricing.plans.pro.name"),
    description: t("pricing.plans.pro.description"),
    priceLabel: t("pricing.plans.pro.priceLabel"),
    priceNote: t("pricing.plans.pro.priceNote"),
    popular: false,
    features: [
      t("pricing.plans.pro.f1"),
      t("pricing.plans.pro.f2"),
      t("pricing.plans.pro.f3"),
      t("pricing.plans.pro.f4"),
      t("pricing.plans.pro.f5"),
    ],
    buttonText: t("pricing.plans.pro.button"),
  },
]);

// Model pricing groups
const modelGroups = computed(() => [
  { category: t("pricing.models.claude") },
  { category: t("pricing.models.openai") },
  { category: t("pricing.models.gemini") },
]);

const flatModels = computed(() => [
  // Claude
  {
    name: "Claude Opus 4.x",
    input: t("pricing.models.prices.claudeOpusInput"),
    output: t("pricing.models.prices.claudeOpusOutput"),
  },
  {
    name: "Claude Sonnet 4.x",
    input: t("pricing.models.prices.claudeSonnetInput"),
    output: t("pricing.models.prices.claudeSonnetOutput"),
  },
  {
    name: "Claude Haiku 4.x",
    input: t("pricing.models.prices.claudeHaikuInput"),
    output: t("pricing.models.prices.claudeHaikuOutput"),
  },
  // OpenAI
  {
    name: "GPT-4.5",
    input: t("pricing.models.prices.gpt45Input"),
    output: t("pricing.models.prices.gpt45Output"),
  },
  {
    name: "GPT-4o",
    input: t("pricing.models.prices.gpt4oInput"),
    output: t("pricing.models.prices.gpt4oOutput"),
  },
  {
    name: "o1 / o3-mini",
    input: t("pricing.models.prices.o1Input"),
    output: t("pricing.models.prices.o1Output"),
  },
  // Gemini
  {
    name: "Gemini 2.5 Pro",
    input: t("pricing.models.prices.geminiProInput"),
    output: t("pricing.models.prices.geminiProOutput"),
  },
  {
    name: "Gemini 2.5 Flash",
    input: t("pricing.models.prices.geminiFlashInput"),
    output: t("pricing.models.prices.geminiFlashOutput"),
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
