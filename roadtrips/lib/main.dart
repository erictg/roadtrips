import 'package:flutter/material.dart';
import 'package:graphql_flutter/graphql_flutter.dart';
import 'package:roadtrips/screens/destination_details.dart';
import 'package:roadtrips/screens/home.dart';
import 'package:roadtrips/screens/search.dart';
import 'package:roadtrips/screens/suggestions.dart';

void main() {
  final _httpLink = HttpLink(
    'https://us-central1-tactical-racer-183815.cloudfunctions.net/DestinationGQL',
  );
  final ValueNotifier<GraphQLClient> client =
      ValueNotifier(GraphQLClient(link: _httpLink, cache: GraphQLCache()));

  runApp(MyApp(client: client));
}

class MyApp extends StatelessWidget {
  final ValueNotifier<GraphQLClient> client;

  const MyApp({super.key, required this.client});

  static List<(String, Widget Function(BuildContext))> routes = [
    HomeScreen.route,
    SearchScreen.route,
    SuggestionsScreen.route,
    DestinationDetailsScreen.route,
  ];

  @override
  Widget build(BuildContext context) {
    return GraphQLProvider(
      client: client,
      child: MaterialApp(
        title: 'Road Tripper',
        theme: ThemeData(
          colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
          useMaterial3: true,
        ),
        initialRoute: '/',
        routes: {for (var route in routes) route.$1: route.$2},
      ),
    );
  }
}
