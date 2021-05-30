import React from 'react';
import {Image} from 'react-native';
import {createBottomTabNavigator} from 'react-navigation-tabs';
import TestView from 'Screens/TestView.js';
import TestView2 from 'Screens/TestView2.js';
import styles from 'Styles/Styles';

export default createBottomTabNavigator(
  {
    TestView: {
      screen: TestView,
      navigationOptions: {
        tabBarIcon: ({focused}) => (
          <Image
            source={focused ? require('icons/facebook_hover.png') : require('icons/facebook.png')}
            style={{width: 10, height: 21}}
          />
        ),
      },
    },
    TestView2: {
      screen: TestView2,
      navigationOptions: {
        tabBarIcon: ({focused}) => (
          <Image
            source={focused ? require('icons/facebook_hover.png') : require('icons/facebook.png')}
            style={{width: 10, height: 21}}
          />
        ),
      },
    },
  },
  {
    tabBarOptions: {
      showLabel: false,
      initialRouteName: 'TestView',
      style: {
        backgroundColor: `${styles.blackColor}`,
      },
    },
  },
);
