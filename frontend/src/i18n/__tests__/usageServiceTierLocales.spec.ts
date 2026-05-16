import { describe, expect, it } from 'vitest'

import en from '../locales/en'
import vi from '../locales/vi'
import zh from '../locales/zh'

describe('usage service tier locale keys', () => {
  it('contains zh labels for service tier tooltip', () => {
    expect(zh.usage.serviceTier).toBe('服务档位')
    expect(zh.usage.serviceTierPriority).toBe('Fast')
    expect(zh.usage.serviceTierFlex).toBe('Flex')
    expect(zh.usage.serviceTierStandard).toBe('Standard')
  })

  it('contains en labels for service tier tooltip', () => {
    expect(en.usage.serviceTier).toBe('Service tier')
    expect(en.usage.serviceTierPriority).toBe('Fast')
    expect(en.usage.serviceTierFlex).toBe('Flex')
    expect(en.usage.serviceTierStandard).toBe('Standard')
  })

  it('contains vi labels for core key usage entries', () => {
    expect(vi.keyUsage.title).toBe('Tra cứu sử dụng API Key')
    expect(vi.keyUsage.dayShort).toBe('Ngày')
    expect(vi.common.autoRefresh.countdown).toBe('Tự động làm mới sau: {seconds} giây')
  })
})
