<template>
  <div
    class="relative min-h-screen overflow-hidden bg-gradient-to-br from-gray-50 via-primary-50/30 to-gray-100 dark:from-dark-950 dark:via-dark-900 dark:to-dark-950"
  >
    <!-- Background Decorations -->
    <div class="pointer-events-none absolute inset-0 overflow-hidden">
      <div
        class="absolute -right-40 -top-40 h-96 w-96 rounded-full bg-primary-400/20 blur-3xl"
      ></div>
      <div
        class="absolute -bottom-40 -left-40 h-96 w-96 rounded-full bg-primary-500/15 blur-3xl"
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

    <main class="relative z-10 px-4 py-8 sm:px-6 lg:px-8">
      <div class="mx-auto max-w-7xl">
        <!-- Hero Section -->
        <section class="mb-12 text-center">
          <p
            class="mb-4 inline-flex items-center gap-2 rounded-full border border-primary-200 bg-white/70 px-4 py-2 text-sm font-medium text-primary-700 shadow-sm backdrop-blur-sm dark:border-primary-800 dark:bg-dark-800/70 dark:text-primary-300"
          >
            <Icon name="book" size="sm" />
            {{ t("docs.badge") }}
          </p>
          <h1
            class="mb-4 text-4xl font-bold text-gray-900 dark:text-white md:text-5xl"
          >
            {{ t("docs.title") }}
          </h1>
          <p
            class="mx-auto mb-8 max-w-3xl text-lg text-gray-600 dark:text-dark-300"
          >
            {{ t("docs.subtitle") }}
          </p>

          <!-- Quick Action Buttons -->
          <div class="flex flex-wrap justify-center gap-4">
            <router-link
              :to="isAuthenticated ? dashboardPath : '/login'"
              class="btn btn-primary px-8 py-3 text-base shadow-lg shadow-primary-500/30"
            >
              {{
                isAuthenticated ? t("home.goToDashboard") : t("home.getStarted")
              }}
              <Icon name="arrowRight" size="md" class="ml-2" />
            </router-link>
            <a
              v-if="docUrl"
              :href="docUrl"
              target="_blank"
              rel="noopener noreferrer"
              class="inline-flex items-center rounded-full border border-gray-300 bg-white px-6 py-3 text-sm font-medium text-gray-700 transition-colors hover:bg-gray-50 dark:border-dark-600 dark:bg-dark-800 dark:text-dark-200 dark:hover:bg-dark-700"
            >
              {{ t("docs.openExternal") }}
              <Icon name="externalLink" size="md" class="ml-2" />
            </a>
          </div>
        </section>

        <!-- Documentation Content -->
        <div
          class="rounded-2xl border border-gray-200/50 bg-white/70 shadow-xl backdrop-blur-sm dark:border-dark-700/50 dark:bg-dark-800/70"
        >
          <!-- Mobile Category Selector -->
          <div class="mb-6 lg:hidden">
            <select
              v-model="activeSection"
              class="w-full rounded-lg border border-gray-300 bg-white px-4 py-3 text-gray-900 dark:border-dark-600 dark:bg-dark-800 dark:text-white"
            >
              <option v-for="cat in categories" :key="cat.id" :value="cat.id">
                {{ cat.label }}
              </option>
            </select>
          </div>

          <div class="flex flex-col lg:flex-row">
            <!-- Sidebar Navigation -->
            <aside
              class="hidden w-64 shrink-0 border-b border-gray-200/50 p-6 dark:border-dark-700/50 lg:block lg:border-b-0 lg:border-r"
            >
              <nav class="sticky top-24 space-y-1">
                <button
                  v-for="cat in categories"
                  :key="cat.id"
                  @click="activeSection = cat.id"
                  :class="[
                    'w-full rounded-lg px-4 py-2 text-left text-sm font-medium transition-colors',
                    activeSection === cat.id
                      ? 'bg-primary-100 text-primary-700 dark:bg-primary-900/30 dark:text-primary-300'
                      : 'text-gray-600 hover:bg-gray-100 dark:text-dark-400 dark:hover:bg-dark-700',
                  ]"
                >
                  {{ cat.label }}
                </button>
              </nav>
            </aside>

            <!-- Main Content -->
            <div class="min-w-0 flex-1 p-6">
              <!-- Introduction -->
              <div
                v-if="activeSection === 'introduction'"
                class="prose dark:prose-invert max-w-none"
              >
                <h1>{{ t("docs.introduction.title") }}</h1>
                <p class="lead">{{ t("docs.introduction.subtitle") }}</p>

                <h2>{{ t("docs.introduction.whyChoose.title") }}</h2>

                <h3>
                  {{ t("docs.introduction.whyChoose.costSavings.title") }}
                </h3>
                <p>
                  {{ t("docs.introduction.whyChoose.costSavings.description") }}
                </p>

                <h3>
                  {{ t("docs.introduction.whyChoose.compatibility.title") }}
                </h3>
                <p>
                  {{
                    t("docs.introduction.whyChoose.compatibility.description")
                  }}
                </p>
                <ul>
                  <li>{{ t("docs.introduction.features.streaming") }}</li>
                  <li>{{ t("docs.introduction.features.vision") }}</li>
                  <li>{{ t("docs.introduction.features.functionCalling") }}</li>
                  <li>{{ t("docs.introduction.features.reasoning") }}</li>
                  <li>{{ t("docs.introduction.features.promptCaching") }}</li>
                </ul>

                <h3>
                  {{ t("docs.introduction.whyChoose.quickIntegration.title") }}
                </h3>
                <p>
                  {{
                    t(
                      "docs.introduction.whyChoose.quickIntegration.description",
                    )
                  }}
                </p>

                <div
                  class="relative group mb-4 rounded-lg bg-gray-900 p-4 overflow-x-auto"
                >
                  <button
                    @click="copyCode"
                    data-copy-id="introduction-urls"
                    class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"
                  >
                    <Icon
                      :name="
                        copiedId === 'introduction-urls' ? 'check' : 'copy'
                      "
                      size="sm"
                    />
                  </button>
                  <p class="text-green-400">
                    OpenAI: https://api.openai.com →
                    https://api.aishopacc.com/v1
                  </p>
                  <p class="text-green-400">
                    Anthropic: https://api.anthropic.com →
                    https://api.aishopacc.com
                  </p>
                </div>

                <h2>{{ t("docs.introduction.howItWorks.title") }}</h2>
                <ol>
                  <li>{{ t("docs.introduction.howItWorks.step1") }}</li>
                  <li>{{ t("docs.introduction.howItWorks.step2") }}</li>
                  <li>{{ t("docs.introduction.howItWorks.step3") }}</li>
                  <li>{{ t("docs.introduction.howItWorks.step4") }}</li>
                  <li>{{ t("docs.introduction.howItWorks.step5") }}</li>
                  <li>{{ t("docs.introduction.howItWorks.step6") }}</li>
                </ol>

                <h2>{{ t("docs.introduction.support.title") }}</h2>
                <ul>
                  <li>
                    {{ t("docs.introduction.support.telegramBot") }}:
                    @aishopacc_bot
                  </li>
                  <li>
                    {{ t("docs.introduction.support.telegramSupport") }}:
                    @aishopacc_support
                  </li>
                  <li>
                    {{ t("docs.introduction.support.email") }}:
                    admin@aishopacc.com
                  </li>
                </ul>
              </div>

              <!-- Purchase & Topup -->
              <div
                v-if="activeSection === 'purchase'"
                class="prose dark:prose-invert max-w-none"
              >
                <h1>{{ t("docs.purchase.title") }}</h1>
                <p>{{ t("docs.purchase.subtitle") }}</p>

                <h2>{{ t("docs.purchase.quickBuy.title") }}</h2>
                <p>{{ t("docs.purchase.quickBuy.description") }}</p>

                <h3>{{ t("docs.purchase.steps.buyKey.title") }}</h3>
                <ol>
                  <li>{{ t("docs.purchase.steps.buyKey.step1") }}</li>
                  <li>{{ t("docs.purchase.steps.buyKey.step2") }}</li>
                  <li>{{ t("docs.purchase.steps.buyKey.step3") }}</li>
                  <li>{{ t("docs.purchase.steps.buyKey.step4") }}</li>
                  <li>{{ t("docs.purchase.steps.buyKey.step5") }}</li>
                  <li>{{ t("docs.purchase.steps.buyKey.step6") }}</li>
                </ol>

                <h3>{{ t("docs.purchase.steps.topup.title") }}</h3>
                <ol>
                  <li>{{ t("docs.purchase.steps.topup.step1") }}</li>
                  <li>{{ t("docs.purchase.steps.topup.step2") }}</li>
                  <li>{{ t("docs.purchase.steps.topup.step3") }}</li>
                  <li>{{ t("docs.purchase.steps.topup.step4") }}</li>
                </ol>

                <h2>{{ t("docs.purchase.groupGuide.title") }}</h2>
                <p>{{ t("docs.purchase.groupGuide.description") }}</p>

                <h3>{{ t("docs.purchase.groupGuide.howToChoose.title") }}</h3>
                <ul>
                  <li>
                    {{ t("docs.purchase.groupGuide.howToChoose.multiModel") }}
                  </li>
                  <li>
                    {{ t("docs.purchase.groupGuide.howToChoose.claude") }}
                  </li>
                  <li>
                    {{ t("docs.purchase.groupGuide.howToChoose.gemini") }}
                  </li>
                  <li>{{ t("docs.purchase.groupGuide.howToChoose.media") }}</li>
                </ul>
              </div>

              <!-- OpenAI Format -->
              <div
                v-if="activeSection === 'openai'"
                class="prose dark:prose-invert max-w-none"
              >
                <h1>{{ t("docs.openai.title") }}</h1>
                <p>{{ t("docs.openai.subtitle") }}</p>

                <h2>{{ t("docs.openai.installation.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="openai-install" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'openai-install' ? 'check' : 'copy'" size="sm" /></button><code>npm install openai
# or
yarn add openai</code></pre>

                <h2>{{ t("docs.openai.setup.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="openai-setup" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'openai-setup' ? 'check' : 'copy'" size="sm" /></button><code>import OpenAI from 'openai';

const openai = new OpenAI({
  apiKey: 'your-aishopacc-api-key',
  baseURL: 'https://api.aishopacc.com/v1',
});</code></pre>

                <h2>{{ t("docs.openai.chatCompletion.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="openai-chat" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'openai-chat' ? 'check' : 'copy'" size="sm" /></button><code>const completion = await openai.chat.completions.create({
  model: 'gpt-4o',
  messages: [
    { role: 'system', content: 'You are a helpful assistant.' },
    { role: 'user', content: 'Hello!' }
  ],
  max_tokens: 1000,
});

console.log(completion.choices[0].message.content);</code></pre>

                <h2>{{ t("docs.openai.streaming.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="openai-stream" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'openai-stream' ? 'check' : 'copy'" size="sm" /></button><code>const stream = await openai.chat.completions.create({
  model: 'gpt-4o',
  messages: [{ role: 'user', content: 'Tell me a story' }],
  stream: true,
});

for await (const chunk of stream) {
  const content = chunk.choices[0]?.delta?.content || '';
  if (content) process.stdout.write(content);
}</code></pre>

                <h2>{{ t("docs.openai.vision.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="openai-vision" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'openai-vision' ? 'check' : 'copy'" size="sm" /></button><code>const completion = await openai.chat.completions.create({
  model: 'gpt-4o',
  messages: [{
    role: 'user',
    content: [
      { type: 'text', text: 'What is in this image?' },
      { type: 'image_url', image_url: { url: 'https://example.com/image.png' } }
    ]
  }],
  max_tokens: 1000,
});</code></pre>

                <h2>{{ t("docs.openai.functionCalling.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="openai-func" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'openai-func' ? 'check' : 'copy'" size="sm" /></button><code>const completion = await openai.chat.completions.create({
  model: 'gpt-4o',
  messages: [{ role: 'user', content: 'What is the weather in Hanoi?' }],
  tools: [{
    type: 'function',
    function: {
      name: 'get_weather',
      description: 'Get weather information',
      parameters: {
        type: 'object',
        properties: { location: { type: 'string' } },
        required: ['location']
      }
    }
  }],
  tool_choice: 'auto',
});</code></pre>

                <h2>{{ t("docs.openai.endpoints.title") }}</h2>
                <table>
                  <thead>
                    <tr>
                      <th>{{ t("docs.openai.endpoints.method") }}</th>
                      <th>{{ t("docs.openai.endpoints.url") }}</th>
                      <th>{{ t("docs.openai.endpoints.description") }}</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>POST</td>
                      <td>/v1/chat/completions</td>
                      <td>{{ t("docs.openai.endpoints.chatCompletion") }}</td>
                    </tr>
                    <tr>
                      <td>POST</td>
                      <td>/v1/images/generations</td>
                      <td>{{ t("docs.openai.endpoints.imageGeneration") }}</td>
                    </tr>
                    <tr>
                      <td>POST</td>
                      <td>/v1/audio/speech</td>
                      <td>{{ t("docs.openai.endpoints.tts") }}</td>
                    </tr>
                    <tr>
                      <td>POST</td>
                      <td>/v1/embeddings</td>
                      <td>{{ t("docs.openai.endpoints.embeddings") }}</td>
                    </tr>
                    <tr>
                      <td>POST</td>
                      <td>/v1/rerank</td>
                      <td>{{ t("docs.openai.endpoints.rerank") }}</td>
                    </tr>
                    <tr>
                      <td>GET</td>
                      <td>/v1/models</td>
                      <td>{{ t("docs.openai.endpoints.listModels") }}</td>
                    </tr>
                  </tbody>
                </table>
              </div>

              <!-- Anthropic Format -->
              <div
                v-if="activeSection === 'anthropic'"
                class="prose dark:prose-invert max-w-none"
              >
                <h1>{{ t("docs.anthropic.title") }}</h1>
                <p>{{ t("docs.anthropic.subtitle") }}</p>

                <h2>{{ t("docs.anthropic.installation.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="anthropic-install" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'anthropic-install' ? 'check' : 'copy'" size="sm" /></button><code>npm install @anthropic-ai/sdk
# or
yarn add @anthropic-ai/sdk</code></pre>

                <h2>{{ t("docs.anthropic.setup.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="anthropic-setup" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'anthropic-setup' ? 'check' : 'copy'" size="sm" /></button><code>import Anthropic from "@anthropic-ai/sdk";

const anthropic = new Anthropic({
  apiKey: "your-aishopacc-api-key",
  baseURL: "https://api.aishopacc.com",
});</code></pre>

                <h2>{{ t("docs.anthropic.messages.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="anthropic-messages" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'anthropic-messages' ? 'check' : 'copy'" size="sm" /></button><code>const message = await anthropic.messages.create({
  model: "claude-3-5-sonnet-20241022",
  max_tokens: 1024,
  messages: [{ role: "user", content: "Hello!" }],
});

console.log(message.content[0].text);</code></pre>

                <h2>{{ t("docs.anthropic.streaming.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="anthropic-stream" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'anthropic-stream' ? 'check' : 'copy'" size="sm" /></button><code>const stream = await anthropic.messages.stream({
  model: "claude-3-5-sonnet-20241022",
  max_tokens: 1024,
  messages: [{ role: "user", content: "Tell me a story" }],
});

for await (const chunk of stream) {
  if (chunk.type === "content_block_delta") {
    process.stdout.write(chunk.delta.text);
  }
}</code></pre>

                <h2>{{ t("docs.anthropic.vision.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="anthropic-vision" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'anthropic-vision' ? 'check' : 'copy'" size="sm" /></button><code>const message = await anthropic.messages.create({
  model: "claude-3-5-sonnet-20241022",
  max_tokens: 1024,
  messages: [{
    role: "user",
    content: [
      { type: "image", source: { type: "url", url: "https://example.com/image.png" } },
      { type: "text", text: "Describe this image" }
    ]
  }],
});</code></pre>

                <h2>{{ t("docs.anthropic.toolUse.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="anthropic-tool" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'anthropic-tool' ? 'check' : 'copy'" size="sm" /></button><code>const message = await anthropic.messages.create({
  model: "claude-3-5-sonnet-20241022",
  max_tokens: 1024,
  tools: [{
    name: "get_weather",
    description: "Get weather information",
    input_schema: {
      type: "object",
      properties: { location: { type: "string" } },
      required: ["location"]
    }
  }],
  messages: [{ role: "user", content: "Weather in Hanoi?" }],
});</code></pre>

                <h2>{{ t("docs.anthropic.apiDetails.title") }}</h2>
                <p>
                  <strong>Endpoint:</strong>
                  <code>POST https://api.aishopacc.com/v1/messages</code>
                </p>
                <p>
                  <strong>Count Tokens:</strong>
                  <code
                    >POST
                    https://api.aishopacc.com/v1/messages/count_tokens</code
                  >
                </p>
              </div>

              <!-- Google GenAI SDK -->
              <div
                v-if="activeSection === 'google-genai'"
                class="prose dark:prose-invert max-w-none"
              >
                <h1>{{ t("docs.googleGenai.title") }}</h1>
                <p>{{ t("docs.googleGenai.subtitle") }}</p>

                <h2>{{ t("docs.googleGenai.installation.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="google-install" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'google-install' ? 'check' : 'copy'" size="sm" /></button><code>npm install @google/genai
# or
yarn add @google/genai</code></pre>

                <h2>{{ t("docs.googleGenai.setup.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="google-setup" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'google-setup' ? 'check' : 'copy'" size="sm" /></button><code>import { GoogleGenAI } from '@google/genai';

const ai = new GoogleGenAI({
  apiKey: 'your-aishopacc-api-key',
  httpOptions: {
    baseUrl: 'https://api.aishopacc.com',
  },
});</code></pre>

                <h2>{{ t("docs.googleGenai.generateContent.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="google-gencontent" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'google-gencontent' ? 'check' : 'copy'" size="sm" /></button><code>const response = await ai.models.generateContent({
  model: 'gemini-2.5-flash',
  contents: 'Hello! Who are you?',
});

console.log(response.text);</code></pre>

                <h2>{{ t("docs.googleGenai.streaming.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="google-stream" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'google-stream' ? 'check' : 'copy'" size="sm" /></button><code>const response = await ai.models.generateContentStream({
  model: 'gemini-2.5-flash',
  contents: 'Tell me a long story',
});

for await (const chunk of response) {
  if (chunk.text) process.stdout.write(chunk.text);
}</code></pre>

                <h2>{{ t("docs.googleGenai.vision.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="google-vision" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'google-vision' ? 'check' : 'copy'" size="sm" /></button><code>const response = await ai.models.generateContent({
  model: 'gemini-2.5-flash',
  contents: [{
    role: 'user',
    parts: [
      { text: 'Describe this image' },
      { fileData: { fileUri: 'https://example.com/image.png', mimeType: 'image/png' } }
    ]
  }],
});</code></pre>

                <h2>{{ t("docs.googleGenai.functionCalling.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="google-func" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'google-func' ? 'check' : 'copy'" size="sm" /></button><code>import { Type } from '@google/genai';

const getWeather = {
  name: 'get_weather',
  description: 'Get weather information',
  parameters: {
    type: Type.OBJECT,
    properties: { location: { type: Type.STRING } },
    required: ['location']
  }
};

const response = await ai.models.generateContent({
  model: 'gemini-2.5-flash',
  contents: 'Weather in Hanoi?',
  config: { tools: [{ functionDeclarations: [getWeather] }] },
});</code></pre>

                <h2>{{ t("docs.googleGenai.endpoints.title") }}</h2>
                <ul>
                  <li>
                    <code>POST /v1beta/models/{model}:generateContent</code>
                  </li>
                  <li>
                    <code
                      >POST /v1beta/models/{model}:streamGenerateContent</code
                    >
                  </li>
                  <li><code>GET /v1beta/models</code></li>
                </ul>
              </div>

              <!-- Claude Code -->
              <div
                v-if="activeSection === 'claude-code'"
                class="prose dark:prose-invert max-w-none"
              >
                <h1>{{ t("docs.claudeCode.title") }}</h1>
                <p>{{ t("docs.claudeCode.subtitle") }}</p>

                <h2>{{ t("docs.claudeCode.installation.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="claude-install" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'claude-install' ? 'check' : 'copy'" size="sm" /></button><code>npm install -g @anthropic-ai/claude-code</code></pre>

                <h2>{{ t("docs.claudeCode.setup.title") }}</h2>
                <p>{{ t("docs.claudeCode.setup.description") }}</p>

                <h3>{{ t("docs.claudeCode.settingsFile.title") }}</h3>
                <p>
                  <strong>macOS / Linux:</strong>
                  <code>~/.claude/settings.json</code>
                </p>
                <p>
                  <strong>Windows:</strong>
                  <code>%USERPROFILE%\.claude\settings.json</code>
                </p>

                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="claude-settings" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'claude-settings' ? 'check' : 'copy'" size="sm" /></button><code>{
  "env": {
    "ANTHROPIC_BASE_URL": "https://api.aishopacc.com",
    "ANTHROPIC_AUTH_TOKEN": "sk-your-token-here"
  }
}</code></pre>

                <h3>{{ t("docs.claudeCode.envVars.title") }}</h3>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="claude-env" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'claude-env' ? 'check' : 'copy'" size="sm" /></button><code># Linux / macOS
export ANTHROPIC_BASE_URL="https://api.aishopacc.com"
export ANTHROPIC_AUTH_TOKEN="sk-your-token-here"

# Windows PowerShell
$env:ANTHROPIC_BASE_URL = "https://api.aishopacc.com"
$env:ANTHROPIC_AUTH_TOKEN = "sk-your-token-here"</code></pre>

                <h2>{{ t("docs.claudeCode.usage.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="claude-usage" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'claude-usage' ? 'check' : 'copy'" size="sm" /></button><code># Interactive mode
claude

# Single shot
claude "Refactor file utils.js"</code></pre>
              </div>

              <!-- Codex CLI -->
              <div
                v-if="activeSection === 'codex'"
                class="prose dark:prose-invert max-w-none"
              >
                <h1>{{ t("docs.codex.title") }}</h1>
                <p>{{ t("docs.codex.subtitle") }}</p>

                <h2>{{ t("docs.codex.installation.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="codex-install" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'codex-install' ? 'check' : 'copy'" size="sm" /></button><code>npm i -g @openai/codex@latest</code></pre>

                <h2>{{ t("docs.codex.configuration.title") }}</h2>
                <h3>config.toml</h3>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="codex-config" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'codex-config' ? 'check' : 'copy'" size="sm" /></button><code>model_provider = "aishopacc"
model = "gpt-5.2"
model_reasoning_effort = "xhigh"
network_access = "enabled"

[model_providers.aishopacc]
name = "aishopacc"
base_url = "https://api.aishopacc.com/v1"
wire_api = "responses"</code></pre>

                <h3>auth.json</h3>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="codex-auth" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'codex-auth' ? 'check' : 'copy'" size="sm" /></button><code>{
  "OPENAI_API_KEY": "sk-your-api-key-here"
}</code></pre>

                <h2>{{ t("docs.codex.usage.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="codex-usage" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'codex-usage' ? 'check' : 'copy'" size="sm" /></button><code># Interactive mode
codex

# Single shot
codex "Refactor file utils.js"

# Select model
codex --model gpt-5 "Explain function calculateTotal"</code></pre>

                <h2>{{ t("docs.codex.envVars.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="codex-env" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'codex-env' ? 'check' : 'copy'" size="sm" /></button><code>export OPENAI_BASE_URL="https://api.aishopacc.com/v1"
export OPENAI_API_KEY="sk-your-api-key-here"</code></pre>
              </div>

              <!-- Cursor IDE -->
              <div
                v-if="activeSection === 'cursor'"
                class="prose dark:prose-invert max-w-none"
              >
                <h1>{{ t("docs.cursor.title") }}</h1>
                <p>{{ t("docs.cursor.subtitle") }}</p>

                <h2>{{ t("docs.cursor.requirements.title") }}</h2>
                <ul>
                  <li>{{ t("docs.cursor.requirements.cursorPro") }}</li>
                  <li>{{ t("docs.cursor.requirements.apiKey") }}</li>
                </ul>

                <h2>{{ t("docs.cursor.configuration.title") }}</h2>
                <ol>
                  <li>{{ t("docs.cursor.configuration.step1") }}</li>
                  <li>{{ t("docs.cursor.configuration.step2") }}</li>
                  <li>{{ t("docs.cursor.configuration.step3") }}</li>
                </ol>

                <h3>{{ t("docs.cursor.openai.title") }}</h3>
                <ul>
                  <li>{{ t("docs.cursor.openai.apiKey") }}</li>
                  <li>
                    {{ t("docs.cursor.openai.baseUrl") }}:
                    <code>https://api.aishopacc.com/cursor</code>
                  </li>
                </ul>

                <h3>{{ t("docs.cursor.anthropic.title") }}</h3>
                <p>{{ t("docs.cursor.anthropic.description") }}</p>

                <h3>{{ t("docs.cursor.google.title") }}</h3>
                <p>{{ t("docs.cursor.google.description") }}</p>

                <h2>{{ t("docs.cursor.usage.title") }}</h2>
                <ul>
                  <li>
                    <code>Cmd/Ctrl + L</code> -
                    {{ t("docs.cursor.usage.chat") }}
                  </li>
                  <li>
                    <code>Cmd/Ctrl + K</code> -
                    {{ t("docs.cursor.usage.inlineEdit") }}
                  </li>
                </ul>
              </div>

              <!-- Cline -->
              <div
                v-if="activeSection === 'cline'"
                class="prose dark:prose-invert max-w-none"
              >
                <h1>{{ t("docs.cline.title") }}</h1>
                <p>{{ t("docs.cline.subtitle") }}</p>

                <h2>{{ t("docs.cline.configuration.title") }}</h2>
                <p>
                  <strong>Base URL:</strong>
                  <code>https://api.aishopacc.com/v1</code>
                </p>

                <h3>{{ t("docs.cline.settings.title") }}</h3>
                <table>
                  <thead>
                    <tr>
                      <th>{{ t("docs.cline.settings.field") }}</th>
                      <th>{{ t("docs.cline.settings.value") }}</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>API Provider</td>
                      <td>OpenAI Compatible</td>
                    </tr>
                    <tr>
                      <td>Base URL</td>
                      <td><code>https://api.aishopacc.com/v1</code></td>
                    </tr>
                    <tr>
                      <td>API Key</td>
                      <td>{{ t("docs.cline.settings.yourKey") }}</td>
                    </tr>
                    <tr>
                      <td>Model ID</td>
                      <td>{{ t("docs.cline.settings.exampleModel") }}</td>
                    </tr>
                  </tbody>
                </table>

                <h3>{{ t("docs.cline.suggestedModels.title") }}</h3>
                <ul>
                  <li><code>gpt-5.2</code></li>
                  <li><code>gpt-5.3-codex</code></li>
                  <li><code>claude-sonnet-4-5</code></li>
                  <li><code>claude-opus-4-5</code></li>
                  <li><code>deepseek-v3.2</code></li>
                  <li><code>qwen3-coder-plus</code></li>
                </ul>

                <h2>{{ t("docs.cline.usage.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="cline-usage" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'cline-usage' ? 'check' : 'copy'" size="sm" /></button><code># Example requests
"Read this project and explain the login flow"
"Write tests for the current payment component"</code></pre>
              </div>

              <!-- OpenCode -->
              <div
                v-if="activeSection === 'opencode'"
                class="prose dark:prose-invert max-w-none"
              >
                <h1>{{ t("docs.opencode.title") }}</h1>
                <p>{{ t("docs.opencode.subtitle") }}</p>

                <h2>{{ t("docs.opencode.installation.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="opencode-install" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'opencode-install' ? 'check' : 'copy'" size="sm" /></button><code>npm install -g opencode
# or
brew install opencode-ai/tap/opencode</code></pre>

                <h2>{{ t("docs.opencode.auth.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="opencode-auth" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'opencode-auth' ? 'check' : 'copy'" size="sm" /></button><code>opencode auth login
# Select "Other" and enter provider id: aishopacc</code></pre>

                <h2>{{ t("docs.opencode.configFile.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="opencode-config" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'opencode-config' ? 'check' : 'copy'" size="sm" /></button><code>{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "aishopacc": {
      "npm": "@ai-sdk/openai-compatible",
      "name": "AISHOPACC",
      "options": {
        "baseURL": "https://api.aishopacc.com/v1"
      },
      "models": {
        "gpt-5.2": { "name": "GPT-5.2" },
        "claude-sonnet-4-5": { "name": "Claude Sonnet 4.5" }
      }
    }
  }
}</code></pre>

                <h2>{{ t("docs.opencode.usage.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="opencode-usage" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'opencode-usage' ? 'check' : 'copy'" size="sm" /></button><code>opencode
/models</code></pre>
              </div>

              <!-- OpenClaw -->
              <div
                v-if="activeSection === 'openclaw'"
                class="prose dark:prose-invert max-w-none"
              >
                <h1>{{ t("docs.openclaw.title") }}</h1>
                <p>{{ t("docs.openclaw.subtitle") }}</p>

                <h2>{{ t("docs.openclaw.endpoints.title") }}</h2>
                <ul>
                  <li>
                    OpenAI-compatible: <code>https://api.aishopacc.com/v1</code>
                  </li>
                  <li>
                    Anthropic-compatible: <code>https://api.aishopacc.com</code>
                  </li>
                  <li>Google GenAI: <code>https://api.aishopacc.com</code></li>
                </ul>

                <h2>{{ t("docs.openclaw.config.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="openclaw-config" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'openclaw-config' ? 'check' : 'copy'" size="sm" /></button><code>{
  "models": {
    "mode": "merge",
    "providers": {
      "aishopacc": {
        "baseUrl": "https://api.aishopacc.com/v1",
        "apiKey": "sk-xxx",
        "api": "openai-completions",
        "models": [
          { "id": "gpt-5", "name": "GPT-5" },
          { "id": "claude-sonnet-4-5", "name": "Claude Sonnet 4.5" }
        ]
      }
    }
  }
}</code></pre>

                <h2>{{ t("docs.openclaw.usage.title") }}</h2>
                <ol>
                  <li>{{ t("docs.openclaw.usage.step1") }}</li>
                  <li>{{ t("docs.openclaw.usage.step2") }}</li>
                  <li>{{ t("docs.openclaw.usage.step3") }}</li>
                  <li>{{ t("docs.openclaw.usage.step4") }}</li>
                  <li>{{ t("docs.openclaw.usage.step5") }}</li>
                </ol>
              </div>

              <!-- Gemini CLI -->
              <div
                v-if="activeSection === 'gemini-cli'"
                class="prose dark:prose-invert max-w-none"
              >
                <h1>{{ t("docs.geminiCli.title") }}</h1>
                <p>{{ t("docs.geminiCli.subtitle") }}</p>

                <h2>{{ t("docs.geminiCli.installation.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="gemini-install" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'gemini-install' ? 'check' : 'copy'" size="sm" /></button><code># npx (no install)
npx @google/gemini-cli

# npm global
npm install -g @google/gemini-cli

# Homebrew
brew install gemini-cli</code></pre>

                <h2>{{ t("docs.geminiCli.configuration.title") }}</h2>
                <h3>~/.gemini/.env</h3>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="gemini-env" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'gemini-env' ? 'check' : 'copy'" size="sm" /></button><code>GEMINI_API_KEY=sk-your-token-here
GOOGLE_GEMINI_BASE_URL=https://api.aishopacc.com</code></pre>

                <h3>{{ t("docs.geminiCli.envVars.title") }}</h3>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="gemini-export" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'gemini-export' ? 'check' : 'copy'" size="sm" /></button><code># Linux / macOS
export GEMINI_API_KEY="sk-your-token-here"
export GOOGLE_GEMINI_BASE_URL="https://api.aishopacc.com"

# Windows PowerShell
$env:GEMINI_API_KEY = "sk-your-token-here"
$env:GOOGLE_GEMINI_BASE_URL = "https://api.aishopacc.com"</code></pre>

                <h2>{{ t("docs.geminiCli.usage.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="gemini-usage" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'gemini-usage' ? 'check' : 'copy'" size="sm" /></button><code># Interactive mode
gemini

# Single shot
gemini "Refactor file utils.js"

# Select model
gemini --model gemini-3-pro "Explain function"</code></pre>
              </div>

              <!-- Nano Banana -->
              <div
                v-if="activeSection === 'nano-banana'"
                class="prose dark:prose-invert max-w-none"
              >
                <h1>{{ t("docs.nanoBanana.title") }}</h1>
                <p>{{ t("docs.nanoBanana.subtitle") }}</p>

                <h2>{{ t("docs.nanoBanana.endpoint.title") }}</h2>
                <p>
                  <strong>Base URL:</strong>
                  <code>https://api.aishopacc.com</code>
                </p>
                <p>
                  <strong>Endpoint:</strong>
                  <code>POST /images/google/generations</code>
                </p>

                <h2>{{ t("docs.nanoBanana.authentication.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="nano-auth" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'nano-auth' ? 'check' : 'copy'" size="sm" /></button><code>Authorization: Bearer &lt;your-api-key&gt;</code></pre>

                <h2>{{ t("docs.nanoBanana.generate.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="nano-generate" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'nano-generate' ? 'check' : 'copy'" size="sm" /></button><code>curl -X POST 'https://api.aishopacc.com/images/google/generations' \
  -H 'Authorization: Bearer &lt;your-api-key&gt;' \
  -H 'Content-Type: application/json' \
  -d '{
    "model": "nano-banana-2",
    "prompt": "A astronaut riding a horse on the moon, cyberpunk style",
    "size": "1:1",
    "imageSize": "1K",
    "format": "png"
  }'</code></pre>

                <h2>{{ t("docs.nanoBanana.parameters.title") }}</h2>
                <table>
                  <thead>
                    <tr>
                      <th>{{ t("docs.nanoBanana.parameters.name") }}</th>
                      <th>{{ t("docs.nanoBanana.parameters.type") }}</th>
                      <th>{{ t("docs.nanoBanana.parameters.required") }}</th>
                      <th>{{ t("docs.nanoBanana.parameters.description") }}</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>model</td>
                      <td>string</td>
                      <td>Yes</td>
                      <td>nano-banana, nano-banana-2, nano-banana-pro</td>
                    </tr>
                    <tr>
                      <td>prompt</td>
                      <td>string</td>
                      <td>Yes</td>
                      <td>Image description</td>
                    </tr>
                    <tr>
                      <td>size</td>
                      <td>string</td>
                      <td>No</td>
                      <td>Aspect ratio (1:1, 16:9, etc.)</td>
                    </tr>
                    <tr>
                      <td>imageSize</td>
                      <td>string</td>
                      <td>No</td>
                      <td>0.5K, 1K, 2K, 4K</td>
                    </tr>
                    <tr>
                      <td>format</td>
                      <td>string</td>
                      <td>No</td>
                      <td>png or jpeg</td>
                    </tr>
                  </tbody>
                </table>

                <h2>{{ t("docs.nanoBanana.models.title") }}</h2>
                <table>
                  <thead>
                    <tr>
                      <th>Model</th>
                      <th>{{ t("docs.nanoBanana.models.refImages") }}</th>
                      <th>{{ t("docs.nanoBanana.models.sizes") }}</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>nano-banana</td>
                      <td>Max 3</td>
                      <td>Fixed</td>
                    </tr>
                    <tr>
                      <td>nano-banana-2</td>
                      <td>Max 5</td>
                      <td>0.5K, 1K, 2K, 4K</td>
                    </tr>
                    <tr>
                      <td>nano-banana-pro</td>
                      <td>Max 5</td>
                      <td>0.5K, 1K, 2K, 4K</td>
                    </tr>
                  </tbody>
                </table>
              </div>

              <!-- Suno AI -->
              <div
                v-if="activeSection === 'suno'"
                class="prose dark:prose-invert max-w-none"
              >
                <h1>{{ t("docs.suno.title") }}</h1>
                <p>{{ t("docs.suno.subtitle") }}</p>

                <h2>{{ t("docs.suno.baseUrl.title") }}</h2>
                <p><code>https://api.aishopacc.com</code></p>

                <h2>{{ t("docs.suno.workflow.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="suno-workflow" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'suno-workflow' ? 'check' : 'copy'" size="sm" /></button><code>1. POST /suno/submit/music  →  Get taskId
                ↓
2. GET  /suno/fetch/{taskId}  →  Poll every 3s
                ↓
3. Get audio_url, image_url from result</code></pre>

                <h2>{{ t("docs.suno.createMusic.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="suno-create" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'suno-create' ? 'check' : 'copy'" size="sm" /></button><code>curl -X POST 'https://api.aishopacc.com/suno/submit/music' \
  -H 'Authorization: Bearer &lt;your-api-key&gt;' \
  -H 'Content-Type: application/json' \
  -d '{
    "mv": "chirp-v4",
    "gpt_description_prompt": "A happy pop song about summer and the beach",
    "title": "Bright Summer",
    "tags": "pop, summer, upbeat"
  }'</code></pre>

                <h2>{{ t("docs.suno.fetchResult.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="suno-fetch" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'suno-fetch' ? 'check' : 'copy'" size="sm" /></button><code>curl 'https://api.aishopacc.com/suno/fetch/{taskId}' \
  -H 'Authorization: Bearer &lt;your-api-key&gt;'</code></pre>

                <h2>{{ t("docs.suno.models.title") }}</h2>
                <table>
                  <thead>
                    <tr>
                      <th>mv value</th>
                      <th>{{ t("docs.suno.models.version") }}</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>chirp-v3.0</td>
                      <td>v3.0</td>
                    </tr>
                    <tr>
                      <td>chirp-v3.5</td>
                      <td>v3.5</td>
                    </tr>
                    <tr>
                      <td>chirp-v4</td>
                      <td>v4.0 (Recommended)</td>
                    </tr>
                    <tr>
                      <td>chirp-fenix</td>
                      <td>v5.5 (Latest)</td>
                    </tr>
                  </tbody>
                </table>

                <h2>{{ t("docs.suno.status.title") }}</h2>
                <ul>
                  <li>
                    <code>SUBMITTED</code> -
                    {{ t("docs.suno.status.submitted") }}
                  </li>
                  <li>
                    <code>QUEUEING</code> - {{ t("docs.suno.status.queueing") }}
                  </li>
                  <li>
                    <code>PROCESSING</code> -
                    {{ t("docs.suno.status.processing") }}
                  </li>
                  <li>
                    <code>SUCCESS</code> - {{ t("docs.suno.status.success") }}
                  </li>
                  <li>
                    <code>FAILED</code> - {{ t("docs.suno.status.failed") }}
                  </li>
                </ul>
              </div>

              <!-- Seller API -->
              <div
                v-if="activeSection === 'seller-api'"
                class="prose dark:prose-invert max-w-none"
              >
                <h1>{{ t("docs.sellerApi.title") }}</h1>
                <p>{{ t("docs.sellerApi.subtitle") }}</p>

                <h2>{{ t("docs.sellerApi.baseUrl.title") }}</h2>
                <p><code>https://api.aishopacc.com/seller</code></p>

                <h2>{{ t("docs.sellerApi.authentication.title") }}</h2>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="seller-auth" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'seller-auth' ? 'check' : 'copy'" size="sm" /></button><code>Authorization: Bearer sk-xxx
# or
x-secret-key: sk-xxx</code></pre>

                <div
                  class="bg-yellow-100 border-l-4 border-yellow-500 p-4 dark:bg-yellow-900/20"
                >
                  <p class="font-medium text-yellow-800 dark:text-yellow-200">
                    {{ t("docs.sellerApi.warning.title") }}
                  </p>
                  <p class="mt-1 text-yellow-700 dark:text-yellow-300">
                    {{ t("docs.sellerApi.warning.message") }}
                  </p>
                </div>

                <h2>{{ t("docs.sellerApi.quickStart.title") }}</h2>
                <h3>{{ t("docs.sellerApi.quickStart.getInfo.title") }}</h3>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="seller-info" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'seller-info' ? 'check' : 'copy'" size="sm" /></button><code>curl -X GET https://api.aishopacc.com/seller \
  -H "Authorization: Bearer sk-your-seller-secret-key"</code></pre>

                <h3>{{ t("docs.sellerApi.quickStart.createKey.title") }}</h3>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="seller-create" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'seller-create' ? 'check' : 'copy'" size="sm" /></button><code>curl -X POST https://api.aishopacc.com/seller/keys \
  -H "Authorization: Bearer sk-your-seller-secret-key" \
  -H "Content-Type: application/json" \
  -d '{"amount": 10, "type": "cheap", "name": "Customer Key"}'</code></pre>

                <h3>{{ t("docs.sellerApi.quickStart.listKeys.title") }}</h3>
                <pre
                  class="relative group bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto"
                ><button @click="copyCode" data-copy-id="seller-list" class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity p-2 rounded-md bg-gray-700 hover:bg-gray-600 text-gray-300"><Icon :name="copiedId === 'seller-list' ? 'check' : 'copy'" size="sm" /></button><code>curl -X GET https://api.aishopacc.com/seller/keys \
  -H "Authorization: Bearer sk-your-seller-secret-key"</code></pre>

                <h2>{{ t("docs.sellerApi.endpoints.title") }}</h2>
                <table>
                  <thead>
                    <tr>
                      <th>Method</th>
                      <th>Endpoint</th>
                      <th>{{ t("docs.sellerApi.endpoints.description") }}</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>GET</td>
                      <td>/seller</td>
                      <td>{{ t("docs.sellerApi.endpoints.getInfo") }}</td>
                    </tr>
                    <tr>
                      <td>GET</td>
                      <td>/seller/keys</td>
                      <td>{{ t("docs.sellerApi.endpoints.listKeys") }}</td>
                    </tr>
                    <tr>
                      <td>POST</td>
                      <td>/seller/keys</td>
                      <td>{{ t("docs.sellerApi.endpoints.createKey") }}</td>
                    </tr>
                    <tr>
                      <td>POST</td>
                      <td>/seller/keys/bulk-create</td>
                      <td>{{ t("docs.sellerApi.endpoints.bulkCreate") }}</td>
                    </tr>
                    <tr>
                      <td>POST</td>
                      <td>/seller/keys/:id/topup</td>
                      <td>{{ t("docs.sellerApi.endpoints.topup") }}</td>
                    </tr>
                    <tr>
                      <td>POST</td>
                      <td>/seller/keys/:id/deduct</td>
                      <td>{{ t("docs.sellerApi.endpoints.deduct") }}</td>
                    </tr>
                    <tr>
                      <td>PUT</td>
                      <td>/seller/keys/:id/status</td>
                      <td>{{ t("docs.sellerApi.endpoints.updateStatus") }}</td>
                    </tr>
                    <tr>
                      <td>DELETE</td>
                      <td>/seller/keys/:id</td>
                      <td>{{ t("docs.sellerApi.endpoints.deleteKey") }}</td>
                    </tr>
                    <tr>
                      <td>GET</td>
                      <td>/seller/stats</td>
                      <td>{{ t("docs.sellerApi.endpoints.stats") }}</td>
                    </tr>
                    <tr>
                      <td>GET</td>
                      <td>/seller/transactions</td>
                      <td>{{ t("docs.sellerApi.endpoints.transactions") }}</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>

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
        <router-link
          to="/home"
          class="text-sm text-gray-500 transition-colors hover:text-gray-700 dark:text-dark-400 dark:hover:text-white"
        >
          {{ t("docs.backHome") }}
        </router-link>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import { useI18n } from "vue-i18n";
import { useAuthStore, useAppStore } from "@/stores";
import Icon from "@/components/icons/Icon.vue";
import PublicNavbar from "@/components/layout/PublicNavbar.vue";

const { t } = useI18n();

const authStore = useAuthStore();
const appStore = useAppStore();

const siteName = computed(
  () =>
    appStore.cachedPublicSettings?.site_name ||
    appStore.siteName ||
    "AISHOPACC",
);
const siteLogo = computed(
  () => appStore.cachedPublicSettings?.site_logo || appStore.siteLogo || "",
);
const docUrl = computed(
  () => appStore.cachedPublicSettings?.doc_url || appStore.docUrl || "",
);

const isDark = ref(document.documentElement.classList.contains("dark"));
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

const activeSection = ref("introduction");
const copiedId = ref<string | null>(null);

function copyCode(event: Event) {
  const button = event.currentTarget as HTMLElement;
  const pre =
    button.parentElement?.querySelector("code") || button.parentElement;
  if (pre) {
    const text = pre.textContent || "";
    navigator.clipboard.writeText(text).then(() => {
      copiedId.value = button.dataset.copyId || null;
      setTimeout(() => {
        copiedId.value = null;
      }, 2000);
    });
  }
}

const categories = computed(() => [
  { id: "introduction", label: t("docs.categories.introduction") },
  { id: "purchase", label: t("docs.categories.purchase") },
  { id: "openai", label: t("docs.categories.openai") },
  { id: "anthropic", label: t("docs.categories.anthropic") },
  { id: "google-genai", label: t("docs.categories.googleGenai") },
  { id: "claude-code", label: t("docs.categories.claudeCode") },
  { id: "codex", label: t("docs.categories.codex") },
  { id: "cursor", label: t("docs.categories.cursor") },
  { id: "cline", label: t("docs.categories.cline") },
  { id: "opencode", label: t("docs.categories.opencode") },
  { id: "openclaw", label: t("docs.categories.openclaw") },
  { id: "gemini-cli", label: t("docs.categories.geminiCli") },
  { id: "nano-banana", label: t("docs.categories.nanoBanana") },
  { id: "suno", label: t("docs.categories.suno") },
  { id: "seller-api", label: t("docs.categories.sellerApi") },
]);

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

onMounted(() => {
  initTheme();
  authStore.checkAuth();

  if (!appStore.publicSettingsLoaded) {
    appStore.fetchPublicSettings();
  }
});
</script>

<style scoped>
.prose h1 {
  @apply mb-4 text-3xl font-bold text-gray-900 dark:text-white;
}
.prose h2 {
  @apply mt-8 mb-4 text-2xl font-semibold text-gray-900 dark:text-white;
}
.prose h3 {
  @apply mt-6 mb-3 text-xl font-semibold text-gray-900 dark:text-white;
}
.prose p {
  @apply mb-4 text-gray-600 dark:text-dark-300;
}
.prose ul,
.prose ol {
  @apply mb-4 list-inside space-y-2 text-gray-600 dark:text-dark-300;
}
.prose li {
  @apply ml-2;
}
.prose code {
  @apply rounded bg-gray-100 px-2 py-1 font-mono text-sm text-primary-600 dark:text-primary-400;
}
.prose code:is(.dark *) {
  background: transparent !important;
}
.prose code:not(:is(.dark *)):not(.prose pre code) {
  background-color: rgb(17 24 39 / var(--tw-bg-opacity, 1)) !important;
}
.prose pre {
  @apply mb-4 rounded-lg bg-gray-900;
}
.prose pre code {
  @apply block bg-transparent p-0 text-gray-100;
}
.prose table {
  @apply mb-4 w-full border-collapse text-sm;
}
.prose th {
  @apply border border-gray-300 bg-gray-100 px-4 py-2 text-left font-semibold dark:border-dark-600 dark:bg-dark-700;
}
.prose td {
  @apply border border-gray-300 px-4 py-2 dark:border-dark-600;
}
</style>
