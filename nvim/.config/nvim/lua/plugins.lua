local lazypath = vim.fn.stdpath("data") .. "/lazy/lazy.nvim"
if not vim.loop.fs_stat(lazypath) then
  vim.fn.system({
    "git",
    "clone",
    "--filter=blob:none",
    "https://github.com/folke/lazy.nvim.git",
    "--branch=stable", -- latest stable release
    lazypath,
  })
end
vim.opt.rtp:prepend(lazypath)
  vim.opt.runtimepath:append("/home/ben/.local/share/nvim/site")
require("lazy").setup({
  { 'catppuccin/nvim', name = 'catppuccin' },
  { 'nvim-tree/nvim-web-devicons' },
  { 'nvim-lualine/lualine.nvim' },
  { 'windwp/nvim-autopairs', event = "InsertEnter" },
  { 'EtiamNullam/deferred-clipboard.nvim' },
  { 'nvim-tree/nvim-tree.lua' },
  { "folke/which-key.nvim" },
  { "nvim-treesitter/nvim-treesitter", build = ":TSUpdate"},
  { 'williamboman/mason.nvim' },
  { 'williamboman/mason-lspconfig.nvim' },
  { 'neovim/nvim-lspconfig' },
  { 'jose-elias-alvarez/null-ls.nvim' },
  { 'hrsh7th/nvim-cmp' },
  { 'hrsh7th/cmp-nvim-lsp' },
  {
	"L3MON4D3/LuaSnip",
	-- follow latest release.
	version = "2.*", -- Replace <CurrentMajor> by the latest released major (first number of latest release)
	-- install jsregexp (optional!).
	build = "make install_jsregexp"
  },
  { 'nvim-lua/plenary.nvim' },
  { 'nvim-telescope/telescope.nvim' },
  { 'Civitasv/cmake-tools.nvim' },
  { 'https://github.com/andweeb/presence.nvim' }
})

require("plugconfig.lualine")
require("plugconfig.nvim-tree")
require("plugconfig.whichkey")
require("plugconfig.lsp")
require("plugconfig.treesitter")
require("plugconfig.cmaketools")
require("plugconfig.presence")

require("nvim-autopairs").setup {}
require('deferred-clipboard').setup {
  fallback = 'unnamedplus', -- or your preferred setting for clipboard
}
