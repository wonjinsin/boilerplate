import {createAppContainer} from 'react-navigation';
import {createStackNavigator} from 'react-navigation-stack';
import MainNavigation from 'Navigations/MainNavigation';
import styles from 'Styles/Styles';
import StackView from 'Screens/StackView';
import {GlobalFn} from 'Fn/GlobalFn';

const NavController = createStackNavigator(
  {
    MainNavigation: {
      screen: MainNavigation,
      navigationOptions: ({navigation}) => ({
        headerShown: false,
      }),
    },
    StackView: {
      screen: StackView,
      navigationOptions: ({navigation}) => ({}),
    }
  },
  {
    mode: 'modal',
    defaultNavigationOptions: {
      headerBackTitle: '',
      headerTintColor: 'white',
      headerStyle: {backgroundColor: styles.blackColor},
      headerBackTitleVisible: null,
    },
  },
);

export default createAppContainer(NavController);
