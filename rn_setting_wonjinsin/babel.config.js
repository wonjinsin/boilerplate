module.exports = {
  presets: ['module:metro-react-native-babel-preset'],
  plugins: [
    [
      'module-resolver',
      {
        root: ['./src'],
        "alias": {
          "icons": "./assets/icons",
          "imgs": "./assets/imgs",
        }
      },
    ],
  ],
};